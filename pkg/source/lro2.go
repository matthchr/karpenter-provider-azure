/*
Portions Copyright (c) Microsoft Corporation.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package source

import (
	"context"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	karpsync "github.com/Azure/karpenter-provider-azure/pkg/sync"
)

type Entry[T any] struct {
	NamespacedName types.NamespacedName
	Data           T
}

// TODO: make this generic?
type LROQueue2 struct {
	lros                    chan Entry[*runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse]]
	MaxConcurrentReconciles int
}

//func (q *LROQueue2) Start(ctx context.Context) {
//	wg := &sync.WaitGroup{}
//
//	go func() {
//		<-ctx.Done()
//		q.lros.ShutDown()
//	}()
//
//	wg.Add(q.MaxConcurrentReconciles)
//	for i := 0; i < q.MaxConcurrentReconciles; i++ {
//		go func() {
//			defer wg.Done()
//			for q.processNextWorkItem(ctx) {
//			}
//		}()
//	}
//
//	<-ctx.Done()
//	c.LogConstructor(nil).Info("Shutdown signal received, waiting for all workers to finish")
//	wg.Wait()
//	c.LogConstructor(nil).Info("All workers finished")
//
//	close(q.lros)
//}
//
//func (q *LROQueue2) processNextWorkItem(ctx context.Context) bool {
//	select {
//	case <-ctx.Done():
//		// Stop working
//		return false
//	default:
//	}
//
//	item := <-q.lros
//
//	// TODO: Process here
//}

func NewLROQueue2() *LROQueue {
	return &LROQueue{
		lros: make(chan event.GenericEvent, 1),
		m:    &karpsync.Map[types.NamespacedName, *runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse]]{},
	}
}

func (q *LROQueue) Add(nodeClaim client.Object, item *runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse]) {
	q.lros <- event.GenericEvent{Object: nodeClaim} // TODO: Unnecessary info?
	nn := types.NamespacedName{Namespace: nodeClaim.GetNamespace(), Name: nodeClaim.GetName()}

	// TODO: What removes items here? A: The controller would need to delete the items from the "provider" when they are done.
	q.m.Store(nn, item) // TODO: this overwrites old info
}

func (q *LROQueue) Source() source.Source {
	return source.Channel(q.lros, handler.Funcs{
		GenericFunc: func(_ context.Context, evt event.GenericEvent, queue workqueue.TypedRateLimitingInterface[reconcile.Request]) {
			nn := types.NamespacedName{Namespace: evt.Object.GetNamespace(), Name: evt.Object.GetName()}
			queue.Add(reconcile.Request{NamespacedName: nn})
		},
	})
}

//func LRO() source.TypedSource[*runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse]] {
//	eventSource := make(chan event.TypedGenericEvent[*runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse]], 1)
//	return source.TypedChannel(eventSource, handler.TypedFuncs[*runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse], *runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse]]{
//		GenericFunc: func(_ context.Context, evt event.TypedGenericEvent[*runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse]], queue workqueue.TypedRateLimitingInterface[*runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse]]) {
//
//			queue.Add()
//		},
//	})
//}

/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	account "github.com/antonum/upjet-rediscloud/internal/controller/account/account"
	database "github.com/antonum/upjet-rediscloud/internal/controller/database/database"
	peering "github.com/antonum/upjet-rediscloud/internal/controller/peering/peering"
	providerconfig "github.com/antonum/upjet-rediscloud/internal/controller/providerconfig"
	subscription "github.com/antonum/upjet-rediscloud/internal/controller/subscription/subscription"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		database.Setup,
		peering.Setup,
		providerconfig.Setup,
		subscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

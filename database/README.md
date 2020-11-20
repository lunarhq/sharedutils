How to use
=============

	import (
		"github.com/lunarhq/sharedutils/database/client"
	)

	db := client.New()

	db.Keys.List()
	db.Keys.Get()
	db.Keys.GetBySecretToken()
	db.Keys.Create( &database.KeyCreateParams{} )
	db.Keys.Update()
	db.Keys.Delete()

	db.Accounts.Create()
	db.Accounts.Update()
	db.Accounts.Delete()
	db.Accounts.List()
	db.Accounts.Get()
	db.Accounts.GetByKey()

	db.PaymentMethods.Create()
	db.PaymentMethods.Update()
	db.PaymentMethods.Delete()
	db.PaymentMethods.List()
	db.PaymentMethods.Get()

	db.UsageRecords.Create()
	db.UsageRecords.Increment()
	db.UsageRecords.Update()
	db.UsageRecords.Delete()
	db.UsageRecords.List()
	db.UsageRecords.Get()

	db.Invoices.Create()
	db.Invoices.Update()
	db.Invoices.Delete()
	db.Invoices.List()
	db.Invoices.Get()

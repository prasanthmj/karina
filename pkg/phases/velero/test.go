package velero

import (
	"github.com/flanksource/commons/console"
	"github.com/flanksource/karina/pkg/platform"
	"github.com/flanksource/kommons"
)

func Test(p *platform.Platform, test *console.TestResults) {
	if p.Velero.IsDisabled() {
		return
	}
	client, err := p.GetClientset()
	if err != nil {
		test.Failf("velero", "Failed to get k8s client %v", err)
		return
	}

	kommons.TestNamespace(client, Namespace, test)

	if !p.E2E {
		return
	}

	if backup, err := CreateBackup(p); err != nil {
		test.Failf("velero", "Failed to create backup: %v", err)
	} else {
		test.Passf("velero", "Backup %s created successfully in %s", backup.Metadata.Name, backup.Status.CompletionTimestamp.Sub(backup.Status.StartTimestamp.Time))
	}
}

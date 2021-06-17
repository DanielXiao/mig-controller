package v1alpha1

// Velero Plugin Annotations
const (
	StageOrFinalMigrationAnnotation = "migrator.run.tanzu.vmware.com/migmigration-type" // (stage|final)
	StageMigration                  = "stage"
	FinalMigration                  = "final"
	PvActionAnnotation              = "migrator.run.tanzu.vmware.com/migrate-type"         // (move|copy)
	PvStorageClassAnnotation        = "migrator.run.tanzu.vmware.com/target-storage-class" // storageClassName
	PvAccessModeAnnotation          = "migrator.run.tanzu.vmware.com/target-access-mode"   // accessMode
	PvCopyMethodAnnotation          = "migrator.run.tanzu.vmware.com/copy-method"          // (snapshot|filesystem)
	QuiesceAnnotation               = "migrator.run.tanzu.vmware.com/migrate-quiesce-pods" // (true|false)
	QuiesceNodeSelector             = "migrator.run.tanzu.vmware.com/quiesceDaemonSet"
	SuspendAnnotation               = "migrator.run.tanzu.vmware.com/preQuiesceSuspend"
	ReplicasAnnotation              = "migrator.run.tanzu.vmware.com/preQuiesceReplicas"
	NodeSelectorAnnotation          = "migrator.run.tanzu.vmware.com/preQuiesceNodeSelector"
	StagePodImageAnnotation         = "migrator.run.tanzu.vmware.com/stage-pod-image"
)

// Restic Annotations
const (
	ResticPvBackupAnnotation = "backup.velero.io/backup-volumes" // comma-separated list of volume names
	ResticPvVerifyAnnotation = "backup.velero.io/verify-volumes" // comma-separated list of volume names
)

// Migration Annotations
const (
	// Disables the internal image copy
	DisableImageCopy = "migrator.run.tanzu.vmware.com/disable-image-copy"
)

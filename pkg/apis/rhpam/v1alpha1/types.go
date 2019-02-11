package v1alpha1

type StatusPhase string

var (
	NoPhase                       StatusPhase
	PhaseInitialized              StatusPhase = "initialized"
	PhaseRealmProvisioned         StatusPhase = "realm provisioned"
	PhasePrepared                 StatusPhase = "prepared"
	PhaseDatabaseInstalled        StatusPhase = "database installed"
	PhaseBusinessCentralInstalled StatusPhase = "business central installed"
	PhaseDatabaseReady            StatusPhase = "database ready"
	PhaseKieServerInstalled       StatusPhase = "kie server installed"
	PhaseComplete                 StatusPhase = "complete"
)
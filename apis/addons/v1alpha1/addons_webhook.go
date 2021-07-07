package v1alpha1

// TODO
// - implement admission.Defaulter + admission.Validator from controller-runtime
// - add cmd/webhook/main.go (separate binary)
// - add build target in Makefile for binary + container (need Dockerfile template?)
// - add manifests to deploy the webhook server as a Deployment? AdmissionWebhook? which k8s resource? Double check these:
//    * Service (livez/readyz)
//    * ServiceMonitor (/metrics ?)
//    * Resource utilization

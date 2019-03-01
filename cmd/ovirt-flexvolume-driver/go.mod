module github.com/feykmeelyahoo/ovirt-openshift-extensions/cmd/ovirt-flexvolume-driver

go 1.12

require (
	github.com/feykmeelyahoo/ovirt-openshift-extensions/internal v0.0.4
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/protobuf v1.3.0 // indirect
	github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/json-iterator/go v1.1.5 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/spf13/viper v1.3.1
	golang.org/x/crypto v0.0.0-20190228161510-8dd112bcdc25 // indirect
	golang.org/x/oauth2 v0.0.0-20190226205417-e64efc72b421 // indirect
	golang.org/x/time v0.0.0-20181108054448-85acf8d2951c // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.0.0-20190227093513-33f4ffca8693
	k8s.io/apimachinery v0.0.0-20190228224630-317ad695e4db
	k8s.io/client-go v2.0.0-alpha.0.0.20190219213553-1f401a01c752+incompatible
	k8s.io/klog v0.2.0 // indirect
	k8s.io/kubernetes v1.13.3
	k8s.io/utils v0.0.0-20190221042446-c2654d5206da // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect
)

replace github.com/feykmeelyahoo/ovirt-openshift-extensions/internal v0.0.4 => ../../internal

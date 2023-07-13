package main

// import (
// 	"gotest.tools/v3/assert"
// 	is "gotest.tools/v3/assert/cmp"
// 	"fmt"
// 	"testing"
// 	"os"
// 	"path/filepath"
// 	"reflect"
// 	"strconv"
// 	"testing"

// 	"github.com/kubernetes/kompose/pkg/kobject"
// 	"github.com/kubernetes/kompose/pkg/loader/compose"
// 	"github.com/kubernetes/kompose/pkg/testutils"
// 	"github.com/pkg/errors"
// 	appsv1 "k8s.io/api/apps/v1"
// 	corev1 "k8s.io/api/core/v1"
// )

// func TestConvertError(t *testing.T) {
// 	randomBuildValue := "random-build"
// 	randomVolumeTypeValue := "random-volume-type"
// 	randomKubernetesControllerValue := "random-controller"
// 	randomKubernetesServiceGroupModeValue := "random-group-mode"
// 	buildConfigValue := string(BUILD_CONFIG)
// 	testCases := []struct {
// 		options      ConvertOptions
// 		errorMessage string
// 	}{
// 		{
// 			options: ConvertOptions{
// 				Build: &randomBuildValue,
// 			},
// 			errorMessage: fmt.Sprintf("unexpected Value for Build field. Possible values are: %v, %v, and %v", string(LOCAL), string(BUILD_CONFIG), string(NONE)),
// 		},
// 		{
// 			options: ConvertOptions{
// 				VolumeType: &randomVolumeTypeValue,
// 			},
// 			errorMessage: fmt.Sprintf("unexpected Value for VolumeType field. Possible values are: %v, %v, %v, %v", string(PVC), string(EMPTYDIR), string(HOSTPATH), string(CONFIGMAP)),
// 		},
// 		{
// 			options: ConvertOptions{
// 				Provider: Kubernetes{
// 					Controller: &randomKubernetesControllerValue,
// 				},
// 			},
// 			errorMessage: fmt.Sprintf("unexpected Value for Kubernetes Controller field. Possible values are: %v, %v, and %v", string(DEPLOYMENT), string(DAEMONSET), string(REPLICATION_CONTROLLER)),
// 		},
// 		{
// 			options: ConvertOptions{
// 				Provider: Kubernetes{
// 					ServiceGroupMode: &randomKubernetesServiceGroupModeValue,
// 				},
// 			},
// 			errorMessage: fmt.Sprintf("unexpected Value for Kubernetes Service Groupe Mode field. Possible values are: %v, %v, ''", string(LABEL), string(VOLUME)),
// 		},
// 		{
// 			options: ConvertOptions{
// 				Provider: Kubernetes{},
// 				Build:    &buildConfigValue,
// 			},
// 			errorMessage: fmt.Sprintf("the build value %v is only supported for Openshift provider", string(BUILD_CONFIG)),
// 		},
// 	}

// 	client, err := NewClient()
// 	assert.Check(t, is.Equal(err, nil))
// 	for _, tc := range testCases {
// 		_, err := client.Convert(tc.options)

// 		assert.Check(t, is.Equal(err.Error(), tc.errorMessage))
// 	}
// }

// func TestConvertWithDefaultOptions(t *testing.T) {
// 	client, err := NewClient(WithErrorOnWarning())
// 	assert.Check(t, is.Equal(err, nil))
// 	objects, err := client.Convert(ConvertOptions{
// 		OutFile: "./testdata/generated/",
// 		InputFiles: []string{
// 			"./testdata/docker-compose.yaml",
// 		},
// 	})
// 	assert.Check(t, is.Equal(err, nil))
// 	for _, object := range objects {
// 		if deployment, ok := object.(*appsv1.Deployment); ok {
// 			assert.Check(t, is.Equal(int(*deployment.Spec.Replicas), 1))
// 		}
// 	}
// }

// func TestNewClientWithOpts(t *testing.T) {
// 	testCases := []struct {
// 		expectedError            error
// 		expectedSuppressWarnings bool
// 		expectedVerbose          bool
// 		expectedErrorOnWarnings  bool
// 		opts                     []Opt
// 	}{
// 		{
// 			expectedError:            nil,
// 			expectedSuppressWarnings: false,
// 			expectedVerbose:          false,
// 			expectedErrorOnWarnings:  false,
// 			opts:                     []Opt{},
// 		},
// 		{
// 			expectedError:            nil,
// 			expectedSuppressWarnings: true,
// 			expectedVerbose:          false,
// 			expectedErrorOnWarnings:  false,
// 			opts:                     []Opt{WithSuppressWarnings()},
// 		},
// 		{
// 			expectedError:            nil,
// 			expectedSuppressWarnings: false,
// 			expectedVerbose:          true,
// 			expectedErrorOnWarnings:  false,
// 			opts:                     []Opt{WithVerboseOutput()},
// 		},
// 		{
// 			expectedError:            nil,
// 			expectedSuppressWarnings: false,
// 			expectedVerbose:          false,
// 			expectedErrorOnWarnings:  true,
// 			opts:                     []Opt{WithErrorOnWarning()},
// 		},
// 		{
// 			expectedError:            nil,
// 			expectedSuppressWarnings: true,
// 			expectedVerbose:          false,
// 			expectedErrorOnWarnings:  true,
// 			opts:                     []Opt{WithErrorOnWarning(), WithSuppressWarnings()},
// 		},
// 	}
// 	for _, tc := range testCases {
// 		client, err := NewClient(tc.opts...)
// 		assert.Check(t, is.Equal(err, tc.expectedError))
// 		assert.Check(t, is.Equal(client.errorOnWarning, tc.expectedErrorOnWarnings))
// 		assert.Check(t, is.Equal(client.verbose, tc.expectedVerbose))
// 		assert.Check(t, is.Equal(client.suppressWarnings, tc.expectedSuppressWarnings))
// 	}
// }

// /*
// Test the creation of a service
// */
// func TestCreateService(t *testing.T) {
// 	// An example service
// 	service := kobject.ServiceConfig{
// 		ContainerName: "name",
// 		Image:         "image",
// 		Environment:   []kobject.EnvVar{{Name: "env", Value: "value"}},
// 		Port:          []kobject.Ports{{HostPort: 123, ContainerPort: 456, Protocol: string(corev1.ProtocolTCP)}},
// 		Command:       []string{"cmd"},
// 		WorkingDir:    "dir",
// 		Args:          []string{"arg1", "arg2"},
// 		VolList:       []string{"/tmp/volume"},
// 		Network:       []string{"network1", "network2"}, // not supported
// 		Labels:        nil,
// 		Annotations:   map[string]string{"abc": "def"},
// 		CPUQuota:      1,                    // not supported
// 		CapAdd:        []string{"cap_add"},  // not supported
// 		CapDrop:       []string{"cap_drop"}, // not supported
// 		Expose:        []string{"expose"},   // not supported
// 		Privileged:    true,
// 		Restart:       "always",
// 	}

// 	// An example object generated via k8s runtime.Objects()
// 	komposeObject := kobject.KomposeObject{
// 		ServiceConfigs: map[string]kobject.ServiceConfig{"app": service},
// 	}
// 	k := Kubernetes{}
// 	_, err := k.Transform(komposeObject, kobject.ConvertOptions{CreateD: true, Replicas: 3})
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "k.Transform failed"))
// 	}

// 	// Test the creation of the service
// 	svc := k.CreateService("foo", service)

// 	if svc.Spec.Ports[0].Port != 123 {
// 		t.Errorf("Expected port 123 upon conversion, actual %d", svc.Spec.Ports[0].Port)
// 	}
// }

// /*
// Test the creation of a service with a memory limit and reservation
// */
// func TestCreateServiceWithMemLimit(t *testing.T) {
// 	// An example service
// 	service := kobject.ServiceConfig{
// 		ContainerName:  "name",
// 		Image:          "image",
// 		Environment:    []kobject.EnvVar{{Name: "env", Value: "value"}},
// 		Port:           []kobject.Ports{{HostPort: 123, ContainerPort: 456, Protocol: string(corev1.ProtocolTCP)}},
// 		Command:        []string{"cmd"},
// 		WorkingDir:     "dir",
// 		Args:           []string{"arg1", "arg2"},
// 		VolList:        []string{"/tmp/volume"},
// 		Network:        []string{"network1", "network2"}, // not supported
// 		Labels:         nil,
// 		Annotations:    map[string]string{"abc": "def"},
// 		CPUQuota:       1,                    // not supported
// 		CapAdd:         []string{"cap_add"},  // not supported
// 		CapDrop:        []string{"cap_drop"}, // not supported
// 		Expose:         []string{"expose"},   // not supported
// 		Privileged:     true,
// 		Restart:        "always",
// 		MemLimit:       1337,
// 		MemReservation: 1338,
// 	}

// 	// An example object generated via k8s runtime.Objects()
// 	komposeObject := kobject.KomposeObject{
// 		ServiceConfigs: map[string]kobject.ServiceConfig{"app": service},
// 	}
// 	k := Kubernetes{}
// 	objects, err := k.Transform(komposeObject, kobject.ConvertOptions{CreateD: true, Replicas: 3})
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "k.Transform failed"))
// 	}

// 	// Retrieve the deployment object and test that it matches the mem value
// 	for _, obj := range objects {
// 		if deploy, ok := obj.(*appsv1.Deployment); ok {
// 			memLimit, _ := deploy.Spec.Template.Spec.Containers[0].Resources.Limits.Memory().AsInt64()
// 			if memLimit != 1337 {
// 				t.Errorf("Expected 1337 for memory limit check, got %v", memLimit)
// 			}
// 			memReservation, _ := deploy.Spec.Template.Spec.Containers[0].Resources.Requests.Memory().AsInt64()
// 			if memReservation != 1338 {
// 				t.Errorf("Expected 1338 for memory reservation check, got %v", memReservation)
// 			}
// 		}
// 	}
// }

// /*
// Test the creation of a service with a cpu limit and reservation
// */
// func TestCreateServiceWithCPULimit(t *testing.T) {
// 	// An example service
// 	service := kobject.ServiceConfig{
// 		ContainerName:  "name",
// 		Image:          "image",
// 		Environment:    []kobject.EnvVar{{Name: "env", Value: "value"}},
// 		Port:           []kobject.Ports{{HostPort: 123, ContainerPort: 456, Protocol: string(corev1.ProtocolTCP)}},
// 		Command:        []string{"cmd"},
// 		WorkingDir:     "dir",
// 		Args:           []string{"arg1", "arg2"},
// 		VolList:        []string{"/tmp/volume"},
// 		Network:        []string{"network1", "network2"}, // not supported
// 		Labels:         nil,
// 		Annotations:    map[string]string{"abc": "def"},
// 		CPUQuota:       1,                    // not supported
// 		CapAdd:         []string{"cap_add"},  // not supported
// 		CapDrop:        []string{"cap_drop"}, // not supported
// 		Expose:         []string{"expose"},   // not supported
// 		Privileged:     true,
// 		Restart:        "always",
// 		CPULimit:       10,
// 		CPUReservation: 1,
// 	}

// 	// An example object generated via k8s runtime.Objects()
// 	komposeObject := kobject.KomposeObject{
// 		ServiceConfigs: map[string]kobject.ServiceConfig{"app": service},
// 	}
// 	k := Kubernetes{}
// 	objects, err := k.Transform(komposeObject, kobject.ConvertOptions{CreateD: true, Replicas: 3})
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "k.Transform failed"))
// 	}

// 	// Retrieve the deployment object and test that it matches the cpu value
// 	for _, obj := range objects {
// 		if deploy, ok := obj.(*appsv1.Deployment); ok {
// 			cpuLimit := deploy.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu().MilliValue()
// 			if cpuLimit != 10 {
// 				t.Errorf("Expected 10 for cpu limit check, got %v", cpuLimit)
// 			}
// 			cpuReservation := deploy.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu().MilliValue()
// 			if cpuReservation != 1 {
// 				t.Errorf("Expected 1 for cpu reservation check, got %v", cpuReservation)
// 			}
// 		}
// 	}
// }

// /*
// Test the creation of a service with a specified user.
// The expected result is that Kompose will set user in PodSpec
// */
// func TestCreateServiceWithServiceUser(t *testing.T) {
// 	// An example service
// 	service := kobject.ServiceConfig{
// 		ContainerName: "name",
// 		Image:         "image",
// 		Environment:   []kobject.EnvVar{{Name: "env", Value: "value"}},
// 		Port:          []kobject.Ports{{HostPort: 123, ContainerPort: 456, Protocol: string(corev1.ProtocolTCP)}},
// 		Command:       []string{"cmd"},
// 		WorkingDir:    "dir",
// 		Args:          []string{"arg1", "arg2"},
// 		VolList:       []string{"/tmp/volume"},
// 		Network:       []string{"network1", "network2"}, // not supported
// 		Labels:        nil,
// 		Annotations:   map[string]string{"kompose.service.type": "nodeport"},
// 		CPUQuota:      1,                    // not supported
// 		CapAdd:        []string{"cap_add"},  // not supported
// 		CapDrop:       []string{"cap_drop"}, // not supported
// 		Expose:        []string{"expose"},   // not supported
// 		Privileged:    true,
// 		Restart:       "always",
// 		User:          "1234",
// 	}

// 	komposeObject := kobject.KomposeObject{
// 		ServiceConfigs: map[string]kobject.ServiceConfig{"app": service},
// 	}
// 	k := Kubernetes{}

// 	objects, err := k.Transform(komposeObject, kobject.ConvertOptions{CreateD: true, Replicas: 1})
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "k.Transform failed"))
// 	}

// 	for _, obj := range objects {
// 		if deploy, ok := obj.(*appsv1.Deployment); ok {
// 			uid := *deploy.Spec.Template.Spec.Containers[0].SecurityContext.RunAsUser
// 			if strconv.FormatInt(uid, 10) != service.User {
// 				t.Errorf("User in ServiceConfig is not matching user in PodSpec")
// 			}
// 		}
// 	}
// }

// func TestTransformWithPid(t *testing.T) {
// 	// An example service
// 	service := kobject.ServiceConfig{
// 		ContainerName: "name",
// 		Image:         "image",
// 		Environment:   []kobject.EnvVar{{Name: "env", Value: "value"}},
// 		Port:          []kobject.Ports{{HostPort: 123, ContainerPort: 456, Protocol: string(corev1.ProtocolTCP)}},
// 		Command:       []string{"cmd"},
// 		WorkingDir:    "dir",
// 		Args:          []string{"arg1", "arg2"},
// 		VolList:       []string{"/tmp/volume"},
// 		Network:       []string{"network1", "network2"},
// 		Restart:       "always",
// 		Pid:           "host",
// 	}

// 	// An example object generated via k8s runtime.Objects()
// 	komposeObject := kobject.KomposeObject{
// 		ServiceConfigs: map[string]kobject.ServiceConfig{"app": service},
// 	}
// 	k := Kubernetes{}
// 	_, err := k.Transform(komposeObject, kobject.ConvertOptions{CreateD: true, Replicas: 3})
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "k.Transform failed"))
// 	}

// 	//for _, obj := range objects {
// 	//	if deploy, ok := obj.(*appsv1.Deployment); ok {
// 	//		hostPid := deploy.Spec.Template.Spec.SecurityContext.HostPID
// 	//		if !hostPid {
// 	//			t.Errorf("Pid in ServiceConfig is not matching HostPID in PodSpec")
// 	//		}
// 	//	}
// 	//}
// }

// func TestTransformWithInvalidPid(t *testing.T) {
// 	// An example service
// 	service := kobject.ServiceConfig{
// 		ContainerName: "name",
// 		Image:         "image",
// 		Environment:   []kobject.EnvVar{{Name: "env", Value: "value"}},
// 		Port:          []kobject.Ports{{HostPort: 123, ContainerPort: 456, Protocol: string(corev1.ProtocolTCP)}},
// 		Command:       []string{"cmd"},
// 		WorkingDir:    "dir",
// 		Args:          []string{"arg1", "arg2"},
// 		VolList:       []string{"/tmp/volume"},
// 		Network:       []string{"network1", "network2"},
// 		Restart:       "always",
// 		Pid:           "badvalue",
// 	}

// 	// An example object generated via k8s runtime.Objects()
// 	komposeObject := kobject.KomposeObject{
// 		ServiceConfigs: map[string]kobject.ServiceConfig{"app": service},
// 	}
// 	k := Kubernetes{}
// 	_, err := k.Transform(komposeObject, kobject.ConvertOptions{CreateD: true, Replicas: 3})
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "k.Transform failed"))
// 	}

// 	//for _, obj := range objects {
// 	//	if deploy, ok := obj.(*appsv1.Deployment); ok {
// 	//		if deploy.Spec.Template.Spec.SecurityContext != nil {
// 	//			hostPid := deploy.Spec.Template.Spec.SecurityContext.HostPID
// 	//			if hostPid {
// 	//				t.Errorf("Pid in ServiceConfig is not matching HostPID in PodSpec")
// 	//			}
// 	//		}
// 	//	}
// 	//}
// }

// func TestIsDir(t *testing.T) {
// 	tempPath := "/tmp/kompose_unit"
// 	tempDir := filepath.Join(tempPath, "i_am_dir")
// 	tempFile := filepath.Join(tempPath, "i_am_file")
// 	tempAbsentDirPath := filepath.Join(tempPath, "i_do_not_exist")

// 	// create directory
// 	err := os.MkdirAll(tempDir, 0744)
// 	if err != nil {
// 		t.Errorf("Unable to create directory: %v", err)
// 	}

// 	// create empty file
// 	f, err := os.Create(tempFile)
// 	if err != nil {
// 		t.Errorf("Unable to create empty file: %v", err)
// 	}
// 	f.Close()

// 	// Check output if directory exists
// 	output, err := isDir(tempDir)
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "isDir failed"))
// 	}
// 	if !output {
// 		t.Errorf("directory %v exists but isDir() returned %v", tempDir, output)
// 	}

// 	// Check output if file is provided
// 	output, err = isDir(tempFile)
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "isDir failed"))
// 	}
// 	if output {
// 		t.Errorf("%v is a file but isDir() returned %v", tempDir, output)
// 	}

// 	// Check output if path does not exist
// 	output, err = isDir(tempAbsentDirPath)
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "isDir failed"))
// 	}
// 	if output {
// 		t.Errorf("Directory %v does not exist, but isDir() returned %v", tempAbsentDirPath, output)
// 	}

// 	// delete temporary directory
// 	err = os.RemoveAll(tempPath)
// 	if err != nil {
// 		t.Errorf("Error removing the temporary directory during cleanup: %v", err)
// 	}
// }

// // TestServiceWithHealthCheck this tests if Headless Service is created for services with HealthCheck.
// func TestServiceWithHealthCheck(t *testing.T) {
// 	testCases := map[string]struct {
// 		service kobject.ServiceConfig
// 	}{
// 		"Exec": {
// 			service: kobject.ServiceConfig{
// 				ContainerName: "name",
// 				Image:         "image",
// 				ServiceType:   "Headless",
// 				HealthChecks: kobject.HealthChecks{
// 					Readiness: kobject.HealthCheck{
// 						Test:        []string{"arg1", "arg2"},
// 						Timeout:     10,
// 						Interval:    5,
// 						Retries:     3,
// 						StartPeriod: 60,
// 					},
// 					Liveness: kobject.HealthCheck{
// 						Test:        []string{"arg1", "arg2"},
// 						Timeout:     11,
// 						Interval:    6,
// 						Retries:     4,
// 						StartPeriod: 61,
// 					},
// 				},
// 			},
// 		},
// 		"HTTPGet": {
// 			service: kobject.ServiceConfig{
// 				ContainerName: "name",
// 				Image:         "image",
// 				ServiceType:   "Headless",
// 				HealthChecks: kobject.HealthChecks{
// 					Readiness: kobject.HealthCheck{
// 						HTTPPath:    "/health",
// 						HTTPPort:    8080,
// 						Timeout:     10,
// 						Interval:    5,
// 						Retries:     3,
// 						StartPeriod: 60,
// 					},
// 					Liveness: kobject.HealthCheck{
// 						HTTPPath:    "/ready",
// 						HTTPPort:    8080,
// 						Timeout:     11,
// 						Interval:    6,
// 						Retries:     4,
// 						StartPeriod: 61,
// 					},
// 				},
// 			},
// 		},
// 		"TCPSocket": {
// 			service: kobject.ServiceConfig{
// 				ContainerName: "name",
// 				Image:         "image",
// 				ServiceType:   "Headless",
// 				HealthChecks: kobject.HealthChecks{
// 					Readiness: kobject.HealthCheck{
// 						TCPPort:     8080,
// 						Timeout:     10,
// 						Interval:    5,
// 						Retries:     3,
// 						StartPeriod: 60,
// 					},
// 					Liveness: kobject.HealthCheck{
// 						TCPPort:     8080,
// 						Timeout:     11,
// 						Interval:    6,
// 						Retries:     4,
// 						StartPeriod: 61,
// 					},
// 				},
// 			},
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		k := Kubernetes{}
// 		komposeObject := kobject.KomposeObject{
// 			ServiceConfigs: map[string]kobject.ServiceConfig{"app": testCase.service},
// 		}
// 		objects, err := k.Transform(komposeObject, kobject.ConvertOptions{CreateD: true, Replicas: 1})
// 		if err != nil {
// 			t.Error(errors.Wrap(err, "k.Transform failed"))
// 		}
// 		if err := testutils.CheckForHealthCheckLivenessAndReadiness(objects); err != nil {
// 			t.Error(err)
// 		}
// 	}
// }

// // TestServiceWithoutPort this tests if Headless Service is created for services without Port.
// func TestServiceWithoutPort(t *testing.T) {
// 	service := kobject.ServiceConfig{
// 		ContainerName: "name",
// 		Image:         "image",
// 		ServiceType:   "Headless",
// 	}

// 	komposeObject := kobject.KomposeObject{
// 		ServiceConfigs: map[string]kobject.ServiceConfig{"app": service},
// 	}
// 	k := Kubernetes{}

// 	objects, err := k.Transform(komposeObject, kobject.ConvertOptions{CreateD: true, Replicas: 1})
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "k.Transform failed"))
// 	}
// 	if err := testutils.CheckForHeadless(objects); err != nil {
// 		t.Error(err)
// 	}
// }

// // Tests if deployment strategy is being set to Recreate when volumes are
// // present
// func TestRecreateStrategyWithVolumesPresent(t *testing.T) {
// 	service := kobject.ServiceConfig{
// 		ContainerName: "name",
// 		Image:         "image",
// 		VolList:       []string{"/tmp/volume"},
// 		Volumes:       []kobject.Volumes{{SvcName: "app", MountPath: "/tmp/volume", PVCName: "app-claim0"}},
// 	}

// 	komposeObject := kobject.KomposeObject{
// 		ServiceConfigs: map[string]kobject.ServiceConfig{"app": service},
// 	}
// 	k := Kubernetes{}

// 	objects, err := k.Transform(komposeObject, kobject.ConvertOptions{CreateD: true, Replicas: 1})
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "k.Transform failed"))
// 	}
// 	for _, obj := range objects {
// 		if deployment, ok := obj.(*appsv1.Deployment); ok {
// 			if deployment.Spec.Strategy.Type != appsv1.RecreateDeploymentStrategyType {
// 				t.Errorf("Expected %v as Strategy Type, got %v",
// 					appsv1.RecreateDeploymentStrategyType,
// 					deployment.Spec.Strategy.Type)
// 			}
// 		}
// 	}
// }

// func TestSortedKeys(t *testing.T) {
// 	service := kobject.ServiceConfig{
// 		ContainerName: "name",
// 		Image:         "image",
// 	}
// 	service1 := kobject.ServiceConfig{
// 		ContainerName: "name",
// 		Image:         "image",
// 	}
// 	c := []string{"a", "b"}

// 	komposeObject := kobject.KomposeObject{
// 		ServiceConfigs: map[string]kobject.ServiceConfig{"b": service, "a": service1},
// 	}
// 	a := SortedKeys(komposeObject)
// 	if !reflect.DeepEqual(a, c) {
// 		t.Logf("Test Fail output should be %s", c)
// 	}
// }

// // test conversion from duration string to seconds *int64
// func TestDurationStrToSecondsInt(t *testing.T) {
// 	testCases := map[string]struct {
// 		in  string
// 		out *int64
// 	}{
// 		"5s":         {in: "5s", out: &[]int64{5}[0]},
// 		"1m30s":      {in: "1m30s", out: &[]int64{90}[0]},
// 		"empty":      {in: "", out: nil},
// 		"onlynumber": {in: "2", out: nil},
// 		"illegal":    {in: "abc", out: nil},
// 	}

// 	for name, test := range testCases {
// 		result, _ := DurationStrToSecondsInt(test.in)
// 		if test.out == nil && result != nil {
// 			t.Errorf("Case '%v' for TestDurationStrToSecondsInt fail, Expected 'nil' , got '%v'", name, *result)
// 		}
// 		if test.out != nil && result == nil {
// 			t.Errorf("Case '%v' for TestDurationStrToSecondsInt fail, Expected '%v' , got 'nil'", name, *test.out)
// 		}
// 		if test.out != nil && result != nil && *test.out != *result {
// 			t.Errorf("Case '%v' for TestDurationStrToSecondsInt fail, Expected '%v' , got '%v'", name, *test.out, *result)
// 		}
// 	}
// }

// func TestServiceWithServiceAccount(t *testing.T) {
// 	assertServiceAccountName := "my-service"

// 	service := kobject.ServiceConfig{
// 		ContainerName: "name",
// 		Image:         "image",
// 		Port:          []kobject.Ports{{HostPort: 55555}},
// 		Labels:        map[string]string{compose.LabelServiceAccountName: assertServiceAccountName},
// 	}

// 	komposeObject := kobject.KomposeObject{
// 		ServiceConfigs: map[string]kobject.ServiceConfig{"app": service},
// 	}
// 	k := Kubernetes{}

// 	objects, err := k.Transform(komposeObject, kobject.ConvertOptions{CreateD: true})
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "k.Transform failed"))
// 	}
// 	for _, obj := range objects {
// 		if deployment, ok := obj.(*appsv1.Deployment); ok {
// 			if deployment.Spec.Template.Spec.ServiceAccountName != assertServiceAccountName {
// 				t.Errorf("Expected %v returned, got %v", assertServiceAccountName, deployment.Spec.Template.Spec.ServiceAccountName)
// 			}
// 		}
// 	}
// }

// func TestCreateServiceWithSpecialName(t *testing.T) {
// 	service := kobject.ServiceConfig{
// 		ContainerName: "front_end",
// 		Image:         "nginx",
// 	}

// 	// An example object generated via k8s runtime.Objects()
// 	komposeObject := kobject.KomposeObject{
// 		ServiceConfigs: map[string]kobject.ServiceConfig{"app": service},
// 	}
// 	k := Kubernetes{}
// 	objects, err := k.Transform(komposeObject, kobject.ConvertOptions{CreateD: true, Replicas: 3})
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "k.Transform failed"))
// 	}
// 	expectedContainerName := "front-end"
// 	for _, obj := range objects {
// 		if deploy, ok := obj.(*appsv1.Deployment); ok {
// 			containerName := deploy.Spec.Template.Spec.Containers[0].Name
// 			if containerName != "front-end" {
// 				t.Errorf("Error while transforming container name. Expected %s Got %s", expectedContainerName, containerName)
// 			}
// 		}
// 	}
// }

// func TestArgsInterpolation(t *testing.T) {
// 	// An example service
// 	service := kobject.ServiceConfig{
// 		ContainerName: "name",
// 		Image:         "image",
// 		Environment:   []kobject.EnvVar{{Name: "PROTOCOL", Value: "https"}, {Name: "DOMAIN", Value: "google.com"}},
// 		Port:          []kobject.Ports{{HostPort: 123, ContainerPort: 456, Protocol: string(corev1.ProtocolTCP)}},
// 		Command:       []string{"curl"},
// 		Args:          []string{"$PROTOCOL://$DOMAIN/"},
// 	}

// 	// An example object generated via k8s runtime.Objects()
// 	komposeObject := kobject.KomposeObject{
// 		ServiceConfigs: map[string]kobject.ServiceConfig{"app": service},
// 	}
// 	k := Kubernetes{}
// 	objects, err := k.Transform(komposeObject, kobject.ConvertOptions{CreateD: true, Replicas: 3})
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "k.Transform failed"))
// 	}

// 	expectedArgs := []string{"$(PROTOCOL)://$(DOMAIN)/"}
// 	for _, obj := range objects {
// 		if deployment, ok := obj.(*appsv1.Deployment); ok {
// 			args := deployment.Spec.Template.Spec.Containers[0].Args[0]
// 			if args != expectedArgs[0] {
// 				t.Errorf("Expected args %v upon conversion, actual %v", expectedArgs, args)
// 			}
// 		}
// 	}
// }

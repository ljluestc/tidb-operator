package member

import (
	"context"
	"github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
	"github.com/pingcap/tidb-operator/pkg/controller"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fakekube "k8s.io/client-go/kubernetes/fake"
	"testing"
)

type fakeDeps struct {
	kubeClient *fakekube.Clientset
}

func (f *fakeDeps) KubeClient() *fakekube.Clientset {
	return f.kubeClient
}

func TestGenerateS3SinkURI(t *testing.T) {
	ns := "testns"
	accessKey := "ak"
	secretKey := "sk"
	kubeClient := fakekube.NewSimpleClientset(&corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "s3-cred",
			Namespace: ns,
		},
		Data: map[string][]byte{
			"access": []byte(accessKey),
			"secret": []byte(secretKey),
		},
	})
	m := &cdcMemberManager{
		deps: &controller.Dependencies{
			KubeClient: kubeClient,
		},
	}
	tc := &v1alpha1.TidbCluster{
		ObjectMeta: metav1.ObjectMeta{Namespace: ns},
		Spec: v1alpha1.TidbClusterSpec{
			CDC: &v1alpha1.CDCSpec{
				Sink: &v1alpha1.SinkSpec{
					Type: "s3",
					Config: v1alpha1.SinkConfig{
						Bucket:                   "bucket",
						Path:                     "cdc-logs",
						AccessKeySecretRef:       &v1alpha1.SecretKeySelector{Name: "s3-cred", Key: "access"},
						SecretAccessKeySecretRef: &v1alpha1.SecretKeySelector{Name: "s3-cred", Key: "secret"},
					},
				},
			},
		},
	}
	uri, err := m.generateS3SinkURI(tc)
	assert.NoError(t, err)
	assert.Equal(t, "s3://bucket/cdc-logs?access-key=ak&secret-access-key=sk", uri)
}

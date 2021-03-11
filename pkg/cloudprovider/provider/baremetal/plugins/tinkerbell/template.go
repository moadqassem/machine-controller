package tinkerbell

import "github.com/tinkerbell/tink/workflow"

func createTemplate(worker, tinkServerAddress, imageRepoAddress string) *workflow.Workflow {
	return &workflow.Workflow{
		Version:       "0.1",
		Name:          "ubuntu_provisioning",
		GlobalTimeout: 6000,
		Tasks: []workflow.Task{
			{
				Name:       "os-installation",
				WorkerAddr: worker,
				Volumes: []string{
					"/dev:/dev",
					"/dev/console:/dev/console",
					"/lib/firmware:/lib/firmware:ro",
				},
				Actions: []workflow.Action{
					{
						Name:    "disk-wipe",
						Image:   "disk-wipe:v1",
						Timeout: 90,
					},
					{
						Name:    "disk-partition",
						Image:   "disk-partition:v1",
						Timeout: 180,
						Environment: map[string]string{
							"MIRROR_HOST": tinkServerAddress,
						},
						Volumes: []string{
							"/statedir:/statedir",
						},
					},
					{
						Name:    "install-root-fs",
						Image:   "install-root-fs:v1",
						Timeout: 600,
						Environment: map[string]string{
							"MIRROR_HOST": imageRepoAddress,
						},
						Volumes: nil,
					},
					{
						Name:    "install-grub",
						Image:   "install-grub:v1",
						Timeout: 600,
						Environment: map[string]string{
							"MIRROR_HOST": imageRepoAddress,
						},
						Volumes: []string{
							"/statedir:/statedir",
						},
					},
				},
			},
		},
	}
}

// TODO(mq): use go templates instead of only formatting the workflow.
var (
	cloudInitWorkflow = `version: '0.1'
name: ubuntu_cloud_init
global_timeout: 6000
tasks:
- name: "cloud-init"
  worker: "%s"
  volumes:
    - /dev:/dev
    - /dev/console:/dev/console
    - /lib/firmware:/lib/firmware:ro
  actions:
  - name: "cloud-init"
    image: cloud-init-start:v1
    timeout: 90`
)

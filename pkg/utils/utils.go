package utils

import (
	"os/exec"

	"github.com/xunzhuo/kube-prow-bot/cmd/kube-prow-bot/config"
	"k8s.io/klog"
)

const GitHubCMD = "gh"

func ExecGitHubCommonCmd(args ...string) error {
	options := append([]string{config.Get().ISSUE_KIND, "-R", config.Get().GH_REPOSITORY}, args...)
	cmd := exec.Command(GitHubCMD, options...)
	cmdOutput, err := cmd.Output()
	if err != nil {
		klog.Error(err)
		return err
	}

	klog.Info(string(cmdOutput))
	return nil
}

func ExecGitHubCmd(args ...string) error {
	cmd := exec.Command(GitHubCMD, args...)
	cmdOutput, err := cmd.Output()
	if err != nil {
		klog.Error(err)
		return err
	}

	klog.Info(string(cmdOutput))
	return nil
}

func ExecGitHubCmdWithOutput(args ...string) (string, error) {
	cmd := exec.Command(GitHubCMD, args...)
	cmdOutput, err := cmd.Output()
	if err != nil {
		klog.Error(err)
		return "", err
	}

	return string(cmdOutput), nil
}
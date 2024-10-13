package vmmanager

import (
	"bytes"
	"c-ademy/internal/vm_manager/commands"
	"c-ademy/internal/vm_manager/language_mappings"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/pkg/errors"
)

type ExecutionResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func ExecuteProgram(
	ctx context.Context,
	executionId string,
	language language_mappings.LanguageKey,
	sourceCode string,
	sourceCodeFileName string,
) (*ExecutionResult, error) {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, errors.Wrap(err, "ExecuteProgram => creating Docker client")
	}
	defer cli.Close()

	dockerImage := language_mappings.LanguageToDockerImage[language]

	reader, err := cli.ImagePull(ctx, dockerImage, image.PullOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "ExecuteProgram => pulling Docker image")
	}
	io.Copy(os.Stdout, reader)

	// Create a temporary directory for the code file
	tempDir, err := os.MkdirTemp("", "code-")
	if err != nil {
		return nil, errors.Wrap(err, "ExecuteProgram => creating temporary directory")
	}
	defer os.RemoveAll(tempDir)

	// Construct the full path
	codePath := filepath.Join(tempDir, sourceCodeFileName)

	// Write the source code to a file
	if err := os.WriteFile(codePath, []byte(sourceCode), 0644); err != nil {
		return nil, errors.Wrap(err, "ExecuteProgram => writing source code to file")
	}

	// Print file contents for debugging
	fileContents, err := os.ReadFile(codePath)
	if err != nil {
		return nil, errors.Wrap(err, "ExecuteProgram => reading source code file")
	}
	fmt.Printf("File contents: %s\n", string(fileContents))

	cmd, err := commands.GenerateDockerCommand(language, "/code/"+sourceCodeFileName)
	if err != nil {
		return nil, errors.Wrap(err, "ExecuteProgram => generating Docker command")
	}

	// Create the container
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:      dockerImage,
		Cmd:        cmd,
		WorkingDir: "/code",
	}, &container.HostConfig{
		Binds: []string{tempDir + ":/code"},
	}, nil, nil, fmt.Sprintf("%s-random-%s", executionId, language))
	if err != nil {
		return nil, errors.Wrap(err, "ExecuteProgram => creating Docker container")
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return nil, errors.Wrap(err, "ExecuteProgram => starting Docker container")
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	var statusCode int64
	select {
	case err := <-errCh:
		if err != nil {
			return nil, errors.Wrap(err, "ExecuteProgram => waiting for container to finish")
		}
	case status := <-statusCh:
		statusCode = status.StatusCode
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return nil, errors.Wrap(err, "ExecuteProgram => fetching container logs")
	}

	var stdout, stderr bytes.Buffer
	stdcopy.StdCopy(&stdout, &stderr, out)

	return &ExecutionResult{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: int(statusCode),
	}, nil
}

package main

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/build"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/opencontainers/image-spec/specs-go/v1"
)

type PodmanImageClient struct {
}

func (p PodmanImageClient) ImageBuild(ctx context.Context, context io.Reader, options build.ImageBuildOptions) (build.ImageBuildResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) BuildCachePrune(ctx context.Context, opts build.CachePruneOptions) (*build.CachePruneReport, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) BuildCancel(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) ImageCreate(ctx context.Context, parentReference string, options image.CreateOptions) (io.ReadCloser, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) ImageImport(ctx context.Context, source image.ImportSource, ref string, options image.ImportOptions) (io.ReadCloser, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) ImageList(ctx context.Context, options image.ListOptions) ([]image.Summary, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) ImagePull(ctx context.Context, ref string, options image.PullOptions) (io.ReadCloser, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) ImagePush(ctx context.Context, ref string, options image.PushOptions) (io.ReadCloser, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) ImageRemove(ctx context.Context, image string, options image.RemoveOptions) ([]image.DeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) ImageSearch(ctx context.Context, term string, options registry.SearchOptions) ([]registry.SearchResult, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) ImageTag(ctx context.Context, image, ref string) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) ImagesPrune(ctx context.Context, pruneFilter filters.Args) (image.PruneReport, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) ImageInspect(ctx context.Context, image string, _ ...client.ImageInspectOption) (image.InspectResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) ImageHistory(ctx context.Context, image string, _ ...client.ImageHistoryOption) ([]image.HistoryResponseItem, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) ImageLoad(ctx context.Context, input io.Reader, _ ...client.ImageLoadOption) (image.LoadResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) ImageSave(ctx context.Context, images []string, _ ...client.ImageSaveOption) (io.ReadCloser, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanImageClient) ImageInspectWithRaw(ctx context.Context, image string) (image.InspectResponse, []byte, error) {
	//TODO implement me
	panic("implement me")
}

type PodmanContainerClient struct {
}

func (p PodmanContainerClient) ContainerAttach(ctx context.Context, container string, options container.AttachOptions) (types.HijackedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerCommit(ctx context.Context, container string, options container.CommitOptions) (container.CommitResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, platform *v1.Platform, containerName string) (container.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerDiff(ctx context.Context, container string) ([]container.FilesystemChange, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerExecAttach(ctx context.Context, execID string, options container.ExecAttachOptions) (types.HijackedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerExecCreate(ctx context.Context, container string, options container.ExecOptions) (container.ExecCreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerExecInspect(ctx context.Context, execID string) (container.ExecInspect, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerExecResize(ctx context.Context, execID string, options container.ResizeOptions) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerExecStart(ctx context.Context, execID string, options container.ExecStartOptions) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerExport(ctx context.Context, container string) (io.ReadCloser, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerInspect(ctx context.Context, container string) (container.InspectResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerInspectWithRaw(ctx context.Context, container string, getSize bool) (container.InspectResponse, []byte, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerKill(ctx context.Context, container, signal string) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerList(ctx context.Context, options container.ListOptions) ([]container.Summary, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerLogs(ctx context.Context, container string, options container.LogsOptions) (io.ReadCloser, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerPause(ctx context.Context, container string) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerRemove(ctx context.Context, container string, options container.RemoveOptions) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerRename(ctx context.Context, container, newContainerName string) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerResize(ctx context.Context, container string, options container.ResizeOptions) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerRestart(ctx context.Context, container string, options container.StopOptions) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerStatPath(ctx context.Context, container, path string) (container.PathStat, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerStats(ctx context.Context, container string, stream bool) (container.StatsResponseReader, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerStatsOneShot(ctx context.Context, container string) (container.StatsResponseReader, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerStart(ctx context.Context, container string, options container.StartOptions) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerStop(ctx context.Context, container string, options container.StopOptions) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerTop(ctx context.Context, container string, arguments []string) (container.TopResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerUnpause(ctx context.Context, container string) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerUpdate(ctx context.Context, container string, updateConfig container.UpdateConfig) (container.UpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainerWait(ctx context.Context, container string, condition container.WaitCondition) (<-chan container.WaitResponse, <-chan error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) CopyFromContainer(ctx context.Context, container, srcPath string) (io.ReadCloser, container.PathStat, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) CopyToContainer(ctx context.Context, container, path string, content io.Reader, options container.CopyToContainerOptions) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanContainerClient) ContainersPrune(ctx context.Context, pruneFilters filters.Args) (container.PruneReport, error) {
	//TODO implement me
	panic("implement me")
}

type PodmanNetworkClient struct {
}

func (p PodmanNetworkClient) NetworkConnect(ctx context.Context, network, container string, config *network.EndpointSettings) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanNetworkClient) NetworkCreate(ctx context.Context, name string, options network.CreateOptions) (network.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanNetworkClient) NetworkDisconnect(ctx context.Context, network, container string, force bool) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanNetworkClient) NetworkInspect(ctx context.Context, network string, options network.InspectOptions) (network.Inspect, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanNetworkClient) NetworkInspectWithRaw(ctx context.Context, network string, options network.InspectOptions) (network.Inspect, []byte, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanNetworkClient) NetworkList(ctx context.Context, options network.ListOptions) ([]network.Summary, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanNetworkClient) NetworkRemove(ctx context.Context, network string) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanNetworkClient) NetworksPrune(ctx context.Context, pruneFilter filters.Args) (network.PruneReport, error) {
	//TODO implement me
	panic("implement me")
}

type PodmanVolumeClient struct {
}

func (p PodmanVolumeClient) VolumeCreate(ctx context.Context, options volume.CreateOptions) (volume.Volume, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanVolumeClient) VolumeInspect(ctx context.Context, volumeID string) (volume.Volume, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanVolumeClient) VolumeInspectWithRaw(ctx context.Context, volumeID string) (volume.Volume, []byte, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanVolumeClient) VolumeList(ctx context.Context, options volume.ListOptions) (volume.ListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanVolumeClient) VolumeRemove(ctx context.Context, volumeID string, force bool) error {
	//TODO implement me
	panic("implement me")
}

func (p PodmanVolumeClient) VolumesPrune(ctx context.Context, pruneFilter filters.Args) (volume.PruneReport, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodmanVolumeClient) VolumeUpdate(ctx context.Context, volumeID string, version swarm.Version, options volume.UpdateOptions) error {
	//TODO implement me
	panic("implement me")
}

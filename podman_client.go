package main

type PodmanClient struct {
	PodmanImageClient
	PodmanContainerClient
	PodmanVolumeClient
	PodmanNetworkClient
}

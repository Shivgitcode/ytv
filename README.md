# YTV - YouTube Viewer

YTV is a command-line tool to stream and download YouTube videos and playlists directly from your terminal.

## Prerequisites

Before using YTV, you need to have the following tools installed on your system:

-   **[yt-dlp](https://github.com/yt-dlp/yt-dlp):** A youtube-dl fork with additional features and fixes.
-   **[mpv](https://mpv.io/):** A free, open source, and cross-platform media player.

## Installation

1.  Install the CLI tool using `go install`:
    ```bash
    go install github.com/Shivgitcode/ytv@latest
    ```
2.  Ensure your Go bin directory is in your system's PATH. This is usually `~/go/bin` on Linux/macOS or `%USERPROFILE%\go\bin` on Windows. If it's not already in your PATH, you can add it by modifying your shell's configuration file (e.g., `.bashrc`, `.zshrc`, `.profile`):
    ```bash
    export PATH=$PATH:$(go env GOPATH)/bin
    ```
    After adding, remember to source your configuration file (e.g., `source ~/.bashrc`) or restart your terminal.

Now you can run `ytv` from any directory.

## Usage

YTV provides three main commands: `stream`, `download`, and `playlist`.

### Stream a Video

To stream a YouTube video, use the `stream` command followed by the video URL.

```bash
ytv stream <video-url>
```

You can also choose the playback speed using the `--speed` flag.

```bash
ytv stream --speed <video-url>
```

This will prompt you to select a playback speed from the available options.

### Download a Video

To download a YouTube video, use the `download` command followed by the video URL.

```bash
ytv download <video-url>
```

By default, the video will be downloaded in the best available quality. You can choose a specific quality using the `--quality` flag.

```bash
ytv download --quality <video-url>
```

This will prompt you to select a quality from the available options (360p, 480p, 720p, 1080p). The downloaded video will be saved in your `Downloads` folder.

### Download a Playlist

To download a YouTube playlist, use the `playlist` command followed by the playlist URL.

```bash
ytv playlist <playlist-url>
```

The videos in the playlist will be downloaded to a folder with the playlist's title inside your `Downloads` folder.

## Flags

-   `--speed`: (stream) Prompts for playback speed selection.
-   `--quality`: (download) Prompts for video quality selection.

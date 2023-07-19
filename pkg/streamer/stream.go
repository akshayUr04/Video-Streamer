package streamer

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Streamer(c *gin.Context) {
	// Fetch video id and playlist name from path parameters
	videoID := c.Param("id")
	playlist := c.Param("playlist")

	playlistData, err := readPlaylistData(videoID, playlist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to read file from server",
			"error":   err.Error(),
		})
		return
	}

	c.Header("Content-Type", "application/vnd.apple.mpegurl") //inform the client that the content type is a hls play list
	c.Header("Content-Disposition", "inline")                 //specify how the client needs to handle the repsonce content
	c.Writer.Write(playlistData)                              //writes the HLS playlist data to the HTTP response body
}

func readPlaylistData(videoID, playlist string) ([]byte, error) {
	// Construct the playlist file path
	playlistPath := fmt.Sprintf("storage/%s/%s", videoID, playlist)
	fmt.Println("---------------read file-------------------")
	// Read the playlist file
	playlistData, err := os.ReadFile(playlistPath)
	if err != nil {
		fmt.Println(err, err.Error())
		return nil, err
	}
	return playlistData, nil
}

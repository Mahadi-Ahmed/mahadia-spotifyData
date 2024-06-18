# mahadia-spotifyData
Use my spotify data somehow

# ToDo:
- [x] Setup containerized psql  
- [x] Make a plan & verify how to perist the data between container restarts & re-creations  
- [] Create schemas for the psql  
- [] Create a script to import spotifyData to psql  

---
# Look into what i can use from Spotify's [web api](https://developer.spotify.com/documentation/web-api)

From the playback data i can use **spotify_track_uri** to request `spotify/track/${id}`
which will return information about that specific track, the track response will contain an artist id which can be used to request `spotify/artists/${artistId}` to get information about genre of the artist

# Analytics & Questions:
- Most played song
    - Genre breakdown of all time
    - Top 10 songs
    - Top 10 artists
- Most skipped song
- Timespan where i usually listen to music
- Most used platform
- Ratio of trackdone vs skipped
- Countries where i've used Spotify
- breakdown of reason_start / reason_end
- ratio of shuffle
- When did i first start to listen on spotify
- total time spent listening
- longest "sessien" of listening to music, in a day
- longest streak of consecutive days of listening to music

- Search for a song, show first time i listened to that song

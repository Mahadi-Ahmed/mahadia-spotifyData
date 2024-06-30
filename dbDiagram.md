# Scratchpad for [db diagrams](https://dbdiagram.io/d)
Table users {
    user_name string [primary key]
}

Table playback {
    id int [primary key]
    user_name string
    ts timestamp
    platform string
    ms_played bigint
    conn_country string
    ip_addr_decrypted string
    user_agent_decrypted string
    reason_start string
    reason_end string
    shuffle boolean
    skipped boolean
    offline boolean
    offline_timestamp bigint
    incognito_mode boolean
    media_type string
}

Table media {
    playback_id int
    media_id string
    media_type string
}

Table track {
    track_id string [primary key]
    track_name string
    artist_name string
    album_name string
    spotify_track_uri string
}

Table podcast {
    podcast_id string [primary key]
    episode_name string
    episode_show_name string
    spotify_episode_uri string
}

Ref: playback.user_name > users.user_name
Ref: media.playback_id > playback.id
Ref: media.media_id > track.track_id
Ref: media.media_id > podcast.podcast_id

<!-- // TODO: Create this table later -->
<!-- // Table Artist {} -->
---

![Db Diagram](dbScetch.png)

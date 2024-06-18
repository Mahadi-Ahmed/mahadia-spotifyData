# Scratchpad for db diagrams
Table Users {
    username string [primary key]
}

Table Playbacks {
    id int [primary key]
    username string
    ts timestamp
    trackId string
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
}

Table Tracks {
    trackId string [primary key]
    track_name string
    album_artist_name string
    album_album_name string
    spotify_track_uri string
    episode_name string
    episode_show_name string
    spotify_episode_uri string
}

Ref: "Users"."username" < "Playbacks"."username"

Ref: "Playbacks"."trackId" < "Tracks"."trackId"

<!-- // TODO: Create this table later -->
<!-- // Table Artist {} -->

Ref: "Users"."username" < "Playbacks"."username"

Ref: "Playbacks"."ts" < "Tracks"."trackId"
---

![Db Diagram](dbScetch.png)

// types.ts
export interface PlaybackData {
  ts: string;
  username: string;
  platform: string;
  ms_played: number;
  conn_country: string;
  ip_addr_decrypted: string;
  user_agent_decrypted: string;
  master_metadata_track_name: string;
  master_metadata_album_artist_name: string;
  master_metadata_album_album_name: string;
  spotify_track_uri: string;
  episode_name: string | null;
  episode_show_name: string | null;
  spotify_episode_uri: string | null;
  reason_start: string;
  reason_end: string;
  shuffle: boolean;
  skipped: boolean | null;
  offline: boolean;
  offline_timestamp: number;
  incognito_mode: boolean;
}

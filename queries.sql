select * from track limit 10;

select playback.track_id, track.track_name, count(playback.track_id) as play_count 
from playback 
join track on playback.track_id = track.track_id 
group by playback.track_id, track.track_name 
order by play_count desc limit 10;


SELECT COUNT(*) AS playback_count
FROM playback;

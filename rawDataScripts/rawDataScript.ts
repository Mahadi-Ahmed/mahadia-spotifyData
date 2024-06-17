// import { data } from '../rawSpotifyData/smallSample.ts'
import type { PlaybackData } from "./types";
import { endsong_0 } from '../rawSpotifyData/MyDataAsTs/endsong_0.ts'
import { endsong_1 } from '../rawSpotifyData/MyDataAsTs/endsong_1.ts'
import { endsong_2 } from '../rawSpotifyData/MyDataAsTs/endsong_2.ts'
import { endsong_3 } from '../rawSpotifyData/MyDataAsTs/endsong_3.ts'
import { endsong_4 } from '../rawSpotifyData/MyDataAsTs/endsong_4.ts'
import { endsong_5 } from '../rawSpotifyData/MyDataAsTs/endsong_5.ts'
import { endsong_6 } from '../rawSpotifyData/MyDataAsTs/endsong_6.ts'
import { endsong_7 } from '../rawSpotifyData/MyDataAsTs/endsong_7.ts'
import { endsong_8 } from '../rawSpotifyData/MyDataAsTs/endsong_8.ts'
import { endsong_9 } from '../rawSpotifyData/MyDataAsTs/endsong_9.ts'

// const playbackData: PlaybackData[] = data
const playbackData0: PlaybackData[] = endsong_0
const playbackData1: PlaybackData[] = endsong_1
const playbackData2: PlaybackData[] = endsong_2
const playbackData3: PlaybackData[] = endsong_3
const playbackData4: PlaybackData[] = endsong_4
const playbackData5: PlaybackData[] = endsong_5
const playbackData6: PlaybackData[] = endsong_6
const playbackData7: PlaybackData[] = endsong_7
const playbackData8: PlaybackData[] = endsong_8
const playbackData9: PlaybackData[] = endsong_9

const endsong0Entries = playbackData0.length
const endsong1Entries = playbackData1.length
const endsong2Entries = playbackData2.length
const endsong3Entries = playbackData3.length
const endsong4Entries = playbackData4.length
const endsong5Entries = playbackData5.length
const endsong6Entries = playbackData6.length
const endsong7Entries = playbackData7.length
const endsong8Entries = playbackData8.length
const endsong9Entries = playbackData9.length

const total = endsong0Entries + endsong1Entries + endsong2Entries + endsong3Entries + endsong4Entries + endsong5Entries + endsong6Entries + endsong7Entries + endsong8Entries + endsong9Entries
console.log(endsong0Entries )
console.log(endsong1Entries )
console.log(endsong2Entries )
console.log(endsong3Entries )
console.log(endsong4Entries )
console.log(endsong5Entries )
console.log(endsong6Entries )
console.log(endsong7Entries )
console.log(endsong8Entries )
console.log(endsong9Entries )
console.log('total entries: ')
console.log(total)

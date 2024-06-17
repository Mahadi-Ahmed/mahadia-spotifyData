import { getBearerToken } from "./auth";

const { access_token } = await getBearerToken()
const headers = {
  'Authorization': `Bearer ${access_token}`,
}

const search = async () => {
  const q = 'track 3clX2NMmjaAHmBjeTSa9vV'
  const type = 'track'
  const url = `${process.env.SPOTIFY_API_URL}/search?q=${q}&type=${type}&market=sv`

  try {
    const response = await fetch(url, { headers })
    const data = await response.json()
    console.log(JSON.stringify(data))
  } catch (error) {
    console.log(error)
  }
}

const track = async (trackId: string) => {
  const url = `${process.env.SPOTIFY_API_URL}/tracks/${trackId}?market=sv`

  try {
    const response = await fetch(url, { headers })
    const data = await response.json()
    console.log(JSON.stringify(data))
  } catch (error) {
    console.log(error)
  }

}

const getArtist = async (artistId: string) => {
  const url = `${process.env.SPOTIFY_API_URL}/artists/${artistId}`

  try {
    const response = await fetch(url, { headers })
    const data = await response.json()
    console.log(JSON.stringify(data))
  } catch (error) {
    console.log(error)
  }
}

// track('3clX2NMmjaAHmBjeTSa9vV')
getArtist('7BMccF0hQFBpP6417k1OtQ')
// search()

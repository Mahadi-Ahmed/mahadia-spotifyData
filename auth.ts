// Client credentials flow

export const getBearerToken = async () => {
  const client = process.env.SPOTIFY_CLIENT
  const secret = process.env.SPOTIFY_SECRET
  const authOptions = {
    method: 'POST',
    headers: {
      'Authorization': 'Basic ' + Buffer.from(client + ':' + secret).toString('base64'),
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    body: new URLSearchParams({
      grant_type: 'client_credentials'
    })
  };

  try {
    const response = await fetch('https://accounts.spotify.com/api/token', authOptions)
    const data = await response.json()
    return data
  } catch (error) {
    console.log(error)
  }
}

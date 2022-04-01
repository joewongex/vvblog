const tokenName = 'api_token'

export const getToken = () => {
  return localStorage.getItem(tokenName)
}

export const setToken = (token) => {
  localStorage.setItem(tokenName, token)
}

export const clearToken = () => {
  localStorage.removeItem(tokenName)
}

export const getUserInfo = () => {
  const token = localStorage.getItem(tokenName)
  if (!token) {
    return null
  }

  // jwt
  try {
    const payload = window.atob(token.split('.', 3)[1])
    return JSON.parse(payload)
  } catch (e) {
    console.error('解析api_token出错', e)
    return null
  }
}
const apiUrl = process.env.EXPO_PUBLIC_API_URL;

export const getOrders= async()=>{
    const res = await fetch(`${apiUrl}/orders`)
    return await res.json()
}
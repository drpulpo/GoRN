import { useEffect, useState } from "react";
import { Order } from "../types";
import { getOrders } from "../api/orders";

export const useOrders=()=>{
    const [orders,setOrders] = useState<Order[]>([])

    async function fetchOrders(){
        setOrders(await getOrders())
    }

    useEffect(()=>{
        fetchOrders()
    },[])
    
    return {
        orders
    }
}
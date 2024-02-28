import { useEffect, useState } from "react";
import { Order } from "../types";
import { getOrders } from "../api/orders";

const initialOrder = {
      "Id": -1,
      "Firstname": "firstname00",
      "Lastname": "lastname00",
      "Phone": "+1 555 555 5555",
      "Email": "email@domain.com",
      "Address1":"address1",
      "Address2":"address2",
      "City": "city00",
      "PostalCode": "72000",
      "State": "ST",
      "Country": "CO",
      "Crayon00": "#000000",
      "Crayon01": "#FFFFFF",
      "Crayon02": "#000000",
      "Crayon03": "#FFFFFF",
      "Crayon04": "#000000",
      "Crayon05": "#FFFFFF",
      "Crayon06": "#000000",
      "Crayon07": "#FFFFFF",
      "Crayon08": "#000000",
      "Crayon09": "#FFFFFF",
      "Crayon10": "#000000",
      "Crayon11": "#FFFFFF",
      "Crayon12": "#000000",
      "Crayon13": "#FFFFFF",
      "Crayon14": "#000000",
      "Crayon15": "#FFFFFF",
      "Crayon16": "#000000",
      "Crayon17": "#FFFFFF",
      "Crayon18": "#000000",
      "Crayon19": "#FFFFFF"
    }

    const exampleOrder = 
        {
          "Id": 1,
          "Firstname": "firstname00",
          "Lastname": "lastname00",
          "Phone": "+1 555 555 5555",
          "Email": "email@domain.com",
          "Address1":"address1",
          "Address2":"address2",
          "City": "city00",
          "PostalCode": "72000",
          "State": "ST",
          "Country": "CO",
          "Crayon00": "#000000",
          "Crayon01": "#FFFFFF",
          "Crayon02": "#000000",
          "Crayon03": "#FFFFFF",
          "Crayon04": "#000000",
          "Crayon05": "#FFFFFF",
          "Crayon06": "#000000",
          "Crayon07": "#FFFFFF",
          "Crayon08": "#000000",
          "Crayon09": "#FFFFFF",
          "Crayon10": "#000000",
          "Crayon11": "#FFFFFF",
          "Crayon12": "#000000",
          "Crayon13": "#FFFFFF",
          "Crayon14": "#000000",
          "Crayon15": "#FFFFFF",
          "Crayon16": "#000000",
          "Crayon17": "#FFFFFF",
          "Crayon18": "#000000",
          "Crayon19": "#FFFFFF"
        }
      

export const useOrders=()=>{
    const [orders,setOrders] = useState<Order[]>([])
    const [selectedOrder,setSelectedOrder] = useState({})

    const fetchOrders = async()=>{
        setOrders(await getOrders())
    }

    useEffect(()=>{
        fetchOrders()
    },[])

    const createOrder=()=>{
        setSelectedOrder(initialOrder)
    }

    const saveOrder=()=>{
        console.log("Saved Selected Order")
        console.log(JSON.stringify(selectedOrder))
    }

    const cancelSaveOrder=()=>{
        setSelectedOrder({})
    }
    
    return {
        orders,
        selectedOrder: selectedOrder,
        createOrder,
        saveOrder,
        cancelSaveOrder
    }
}
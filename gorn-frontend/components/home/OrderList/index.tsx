import { Order } from "../../../types"



export const OrderList = ({ orders }) => {
    return (
        <section>{JSON.stringify(orders)}</section>
    )
}
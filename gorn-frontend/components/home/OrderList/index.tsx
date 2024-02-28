import { Text, View, StyleSheet } from 'react-native'

export const OrderList = ({ orders }) => {
    return (
        <View style={styles.container}>
            <Text>
                TODO: OrdersList || OrderTracking Component
            </Text>
            <Text>
                {JSON.stringify(orders)}
            </Text>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        width: '100%',
        display: 'flex',
        flex: 1,
        backgroundColor: '#ffffff',
        justifyContent: 'space-between',
        padding: 10,
        height: 100
    }
})
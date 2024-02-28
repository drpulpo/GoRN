import { View, Text, Button, StyleSheet, TextInput } from "react-native"

export const OrderForm = ({ selectedOrder, onSave, onCancel }) => {

    return (
        <View style={styles.container}>
            <View style={styles.cancelSaveRow}>
                <Button title="Cancel" color={styles.button.backgroundColor} onPress={() => { onCancel() }} />
                <Button title="Save" color={styles.button.backgroundColor} onPress={() => { onSave() }} />
            </View>
            <View style={styles.nameRow}>
                <TextInput style={styles.firstName} value={selectedOrder.Firstname} placeholder="First Name" />
                <TextInput style={styles.lastName} value={selectedOrder.Lastname} placeholder="Last Name" />
            </View>
            <View style={styles.nameRow}>
                <TextInput style={styles.phone} value={selectedOrder.Phone} placeholder="Phone" />
                <TextInput style={styles.email} value={selectedOrder.Email} placeholder="Email" />
            </View>
            <View style={styles.nameRow}>
                <Text>Todo: Rest of the Fields</Text>
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#FFFFFF',
        padding: 5,
        gap: 25
    },
    cancelSaveRow: {
        flexDirection: 'row',
        justifyContent: 'space-between'
    },
    button: {
        backgroundColor: '#21d321'
    },
    nameRow: {
        height: 50,
        flexDirection: 'row',
        flex: 1,
        gap: 25
    },
    firstName: {
        flex: 1
    },
    lastName: {
        flex: 1
    },
    phone: {
        flex: 1
    },
    email: {
        flex: 1
    }
})
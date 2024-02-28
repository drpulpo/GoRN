import { Text, View, StyleSheet } from 'react-native'

export const Footer = () => {
    return (
        <View style={styles.container}>
            <Text style={styles.about}>David Nuñez 2024</Text>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        width: '100%',
        display: 'flex',
        flexDirection: 'row',
        height: 50,
        backgroundColor: '#21d321',
        color: '#ffffff',
        justifyContent: 'center',
        padding: 10
    },
    about: {
        color: '#FFFFFF',
        fontSize: 18
    }
})
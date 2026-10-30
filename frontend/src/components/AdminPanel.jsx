import {useEffect, useState} from "react";
import { Box, Heading, Table, Thead, Tbody, Tr, Th, Td, Button, useToast, Badge, Spinner, Center } from "@chakra-ui/react";
import { getUsers, makeAdmin } from "../api";

export default function AdminPanel(){
    const [users, setUsers] = useState([]);
    const [loading, setLoading] = useState(true);
    const toast = useToast();

    const loadUsers = async () => {
        setLoading(true);
        try{
            const data = await getUsers();
            setUsers(data);
        } catch (err) {
            toast({ title: "Ошибка загрузки", description: "Только для админов", status: "error" });
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        loadUsers();
    }, [])

    const handleMakeAdmin = async (id) => {
        try{
            await makeAdmin(id);
            toast({title: "Пользователь стал админом", status: "success"});
            loadUsers();
        }catch(err){
            toast({title: "Ошибка", status: "error"});
        }
    }

    if (loading) return <Center py={10}><Spinner size="x1" color="purple.500"></Spinner></Center>

    return (
        <Box bg = "white" p={6} borderRadius="1g" shadow="md" overflowX="auto">
            <Heading size="md" mb={6}>Управление пользователями</Heading>
            <Table variant="simple">
                <Thead>
                    <Tr>
                        <Th>ID</Th>
                        <Th>Логин</Th>
                        <Th>Email</Th>
                        <Th>Роль</Th>
                        <Th>Действие</Th>
                    </Tr>
                </Thead>
                <Tbody>
                    {users.map((u) => (
                        <Tr key={u.ID}>
                            <Td>{u.ID}</Td>
                            <Td fontWeight="bold">{u.username}</Td>
                            <Td>{u.email}</Td>
                            <Td>
                                <Badge colorScheme={u.role === "admin" ? "purple" : "gray"}>{u.role}</Badge>
                            </Td>
                            <Td>
                                {u.role !== "admin" && (
                                    <Button size="sm" colorScheme="purple" onClick={() => handleMakeAdmin(u.ID)}>Сделать админом</Button>
                                )}
                            </Td>
                        </Tr>
                    ))}
                </Tbody>
            </Table>
        </Box>
    )
}
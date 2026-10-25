import { useState } from "react";
import {
    Box, Button, FormControl, FormLabel, Input, VStack, Heading, Text, useToast, Center , Link 
} from "@chakra-ui/react";
import { login, setAccessToken } from "../api";

export default function Login({ onLoginSuccess }) {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [loading, setLoading] = useState(false);
    const [isRegister, setIsRegister] = useState(false); 
    const toast = useToast();

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        try {
            if (isRegister){
                const response = await fetch("http://localhost:3000/api/v1/register", {
                    method: "POST",
                    headers: {"Content-Type": "applicatiob/json"},
                    body: JSON.stringify({username, password})
                })
            

            if (!response.ok) throw new Error("Registration failed");

            toast({title: "Успешная регистрация", description: "Теперь войдите", status: "success"});
            setIsRegister(false);
            }else{
                const data = await login(username, password);
                setAccessToken(data.access_token);
                localStorage.setItem("role", data.role);
                toast({title: "Успешный вход", status: "success"});
                onLoginSuccess();
        } }
        catch(err){
            toast({title: "Ощибка", description: "Проверьте данные", status: "error"});
        } finally {
            setLoading(false);
        }
    };

    return (
    <Center minH="100vh" bg="gray.50">
      <Box bg="white" p={8} borderRadius="lg" shadow="md" w="100%" maxW="400px">
        <VStack as="form" onSubmit={handleSubmit} spacing={4}>
          <Heading size="lg" color="purple.600">
            {isRegister ? "Регистрация" : "Вход"}
          </Heading>
          
          <FormControl isRequired>
            <FormLabel>Логин</FormLabel>
            <Input value={username} onChange={(e) => setUsername(e.target.value)} />
          </FormControl>
          
          <FormControl isRequired>
            <FormLabel>Пароль</FormLabel>
            <Input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
          </FormControl>
          
          <Button type="submit" colorScheme="purple" width="100%" isLoading={loading}>
            {isRegister ? "Создать аккаунт" : "Войти"}
          </Button>

          <Text fontSize="sm" color="gray.500" mt={2}>
            {isRegister ? "Уже есть аккаунт?" : "Нет аккаунта?"}{" "}
            <Link color="purple.500" onClick={() => setIsRegister(!isRegister)} fontWeight="bold">
              {isRegister ? "Войти" : "Зарегистрироваться"}
            </Link>
          </Text>
        </VStack>
      </Box>
    </Center>
    );
}
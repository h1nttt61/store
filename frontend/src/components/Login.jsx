import { useState } from "react";
import { 
  Box, Button, FormControl, FormLabel, Input, VStack, Heading, 
  Text, useToast, Center, Link, PinInput, PinInputField, HStack 
} from "@chakra-ui/react";
import { login, setAccessToken, registerUser, loginWithCode } from "../api";

export default function Login({ onLoginSuccess }) {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [code, setCode] = useState("");
  const [step, setStep] = useState(1);
  const [loading, setLoading] = useState(false);
  const toast = useToast();

  const handleLogin = async (e) => {
    e.preventDefault();
    setLoading(true);
    try {
      const data = await login(username, password);
      
      if (data.require_code) {
        setStep(2);
        toast({ title: "Код отправлен на вашу почту", status: "info" });
      } else {
        setAccessToken(data.access_token);
        localStorage.setItem("role", data.role);
        onLoginSuccess();
      }
    } catch (err) {
      toast({ title: "Неверный логин или пароль", status: "error" });
    } finally {
      setLoading(false);
    }
  };

  const handleVerifyCode = async () => {
    setLoading(true);
    try {
      const data = await loginWithCode(username, code);
      setAccessToken(data.access_token);
      localStorage.setItem("role", data.role);
      toast({ title: "Успешный вход!", status: "success" });
      onLoginSuccess();
    } catch (err) {
      toast({ title: "Неверный код", status: "error" });
    } finally {
      setLoading(false);
    }
  };

  const handleRegister = async (e) => {
    e.preventDefault();
    setLoading(true);
    try {
      await registerUser(username, email, password);
      toast({ title: "Аккаунт создан! Теперь войдите.", status: "success" });
      setStep(1);
    } catch (err) {
      toast({ title: err.response?.data?.error || "Ошибка регистрации", status: "error" });
    } finally {
      setLoading(false);
    }
  };

  return (
    <Center minH="100vh" bg="gray.50">
      <Box bg="white" p={8} borderRadius="lg" shadow="xl" w="100%" maxW="400px">
        
        {step === 1 && (
          <VStack as="form" onSubmit={handleLogin} spacing={4}>
            <Heading size="lg" color="purple.600">Вход</Heading>
            <FormControl isRequired>
              <FormLabel>Логин</FormLabel>
              <Input value={username} onChange={(e) => setUsername(e.target.value)} />
            </FormControl>
            <FormControl isRequired>
              <FormLabel>Пароль</FormLabel>
              <Input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
            </FormControl>
            <Button type="submit" colorScheme="purple" width="100%" isLoading={loading}>
              Войти
            </Button>
            <Text fontSize="sm" color="gray.500">
              Нет аккаунта?{" "}
              <Link color="purple.500" onClick={() => setStep(3)} fontWeight="bold">
                Зарегистрироваться
              </Link>
            </Text>
          </VStack>
        )}

        {step === 2 && (
          <VStack spacing={6}>
            <Heading size="lg" color="purple.600">Введите код</Heading>
            <Text color="gray.500" textAlign="center">
              Мы отправили 6-значный код на вашу почту.
            </Text>
            <HStack>
              <PinInput size="lg" otp value={code} onChange={setCode}>
                <PinInputField /><PinInputField /><PinInputField />
                <PinInputField /><PinInputField /><PinInputField />
              </PinInput>
            </HStack>
            <Button 
              colorScheme="purple" width="100%" onClick={handleVerifyCode} 
              isLoading={loading} isDisabled={code.length !== 6}
            >
              Подтвердить
            </Button>
            <Button variant="ghost" size="sm" onClick={() => setStep(1)}>← Назад</Button>
          </VStack>
        )}

        {step === 3 && (
          <VStack as="form" onSubmit={handleRegister} spacing={4}>
            <Heading size="lg" color="purple.600">Регистрация</Heading>
            <FormControl isRequired>
              <FormLabel>Логин</FormLabel>
              <Input value={username} onChange={(e) => setUsername(e.target.value)} />
            </FormControl>
            <FormControl isRequired>
              <FormLabel>Email</FormLabel>
              <Input type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
            </FormControl>
            <FormControl isRequired>
              <FormLabel>Пароль</FormLabel>
              <Input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
            </FormControl>
            <Button type="submit" colorScheme="purple" width="100%" isLoading={loading}>
              Создать аккаунт
            </Button>
            <Button variant="ghost" size="sm" onClick={() => setStep(1)}>← Уже есть аккаунт</Button>
          </VStack>
        )}
      </Box>
    </Center>
  );
}
import { useEffect, useState } from "react";
import {
  Box, Container, Heading, SimpleGrid, Tabs, TabList, Tab, TabPanels, TabPanel,
  Spinner, Center, Text, useToast, Button, HStack, Flex
} from "@chakra-ui/react";
import { getKittens, getAccessToken, logout, refreshToken } from "./api"; 
import KittenCard from "./components/KittenCard";
import AddKittenForm from "./components/AddKittenForm";
import Login from "./components/Login"; 

export default function App() {
  const [kittens, setKittens] = useState([]);
  const [loading, setLoading] = useState(true);
  const [page, setPage] = useState(1);
  
  const [isAuthenticated, setIsAuthenticated] = useState(!!getAccessToken());
  const [isCheckingAuth, setIsCheckingAuth] = useState(true); // Новый стейт: проверяем токен
  const [userRole, setUserRole] = useState(localStorage.getItem("role") || "guest");
  
  const [showLogin, setShowLogin] = useState(false);
  const toast = useToast();

  useEffect(() => {
    const checkAuth = async () => {
      const token = getAccessToken();
      if (!token) {
        try {
          const data = await refreshToken(); 
          if (data.access_token) {
            setIsAuthenticated(true);
          }
        } catch (err) {
          console.log("Нет валидной сессии", err);
          setIsAuthenticated(false);
        }
      }
      setIsCheckingAuth(false);
    };
    checkAuth();
  }, []);

  const load = async () => {
    setLoading(true);
    try {
      const res = await getKittens(page, 6);
      setKittens(res.data);
    } catch (err) {
      console.error(err);
      toast({ title: "Не удалось загрузить котят", status: "error" });
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (isAuthenticated) load();
  }, [page, isAuthenticated]);

  const handleLogout = () => {
    logout();
    setIsAuthenticated(false);
    setUserRole("guest");
    setShowLogin(false);
  };

  const handleLoginSuccess = () => {
    setIsAuthenticated(true);
    setUserRole(localStorage.getItem("role"));
    setShowLogin(false);
  };

  if (isCheckingAuth) {
    return <Center minH="100vh"><Spinner size="xl" color="purple.500" /></Center>;
  }

  if (showLogin) {
    return (
      <Box minH="100vh" bg="gray.50">
        <Button mt={4} ml={4} onClick={() => setShowLogin(false)}>← Назад в каталог</Button>
        <Login onLoginSuccess={handleLoginSuccess} />
      </Box>
    );
  }

  const isAdmin = userRole === "admin";

  return (
      <Box minH="100vh" bg="gray.50">
        <Box bg="purple.600" color="white" py={6} mb={8}>
          <Container maxW="container.xl">
            <Flex justify="space-between" align="center">
              <Box>
                <Heading size="lg">🐱 Kittens Shop · EU</Heading>
                <Text opacity={0.85} mt={1}>Твой проект по продаже котят</Text>
              </Box>
              {isAuthenticated ? (
                <Button colorScheme="whiteAlpha" size="sm" onClick={handleLogout}>Выйти</Button>
              ) : (
                <Button colorScheme="whiteAlpha" size="sm" onClick={() => setShowLogin(true)}>Войти</Button>
              )}
            </Flex>
          </Container>
        </Box>

        <Container maxW="container.xl" pb={16}>
          <Tabs colorScheme="purple" variant="enclosed">
            <TabList>
              <Tab>Каталог</Tab>
              {isAdmin && <Tab>Добавить</Tab>}
            </TabList>

            <TabPanels>
              <TabPanel px={0}>
                {loading ? (
                    <Center py={20}><Spinner size="xl" color="purple.500" /></Center>
                ) : kittens.length === 0 ? (
                    <Text textAlign="center" py={20} color="gray.500">Пока нет доступных котят</Text>
                ) : (
                    <>
                      <SimpleGrid columns={{ base: 1, md: 2, lg: 3 }} spacing={6}>
                        {kittens.map((k) => (
                            <KittenCard key={k.ID} kitten={k} onDetails={(kitten) => toast({ title: kitten.name, description: kitten.description || "Без описания", status: "info" })} />
                        ))}
                      </SimpleGrid>
                      <Flex justify="center" mt={8}>
                        <HStack>
                          <Button onClick={() => setPage(page - 1)} isDisabled={page === 1}>Назад</Button>
                          <Text fontWeight="bold" alignSelf="center" mx={4}>Страница {page}</Text>
                          <Button onClick={() => setPage(page + 1)} isDisabled={kittens.length < 6}>Вперед</Button>
                        </HStack>
                      </Flex>
                    </>
                )}
              </TabPanel>

              {isAdmin && (
                <TabPanel px={0}>
                  <AddKittenForm onCreated={() => { setPage(1); load(); }} />
                </TabPanel>
              )}
            </TabPanels>
          </Tabs>
        </Container>
      </Box>
  );
}
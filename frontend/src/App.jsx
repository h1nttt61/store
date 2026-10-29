import { useEffect, useState } from "react";
import {
  Box, Container, Heading, SimpleGrid, Tabs, TabList, Tab, TabPanels, TabPanel,
  Spinner, Center, Text, useToast, Button, HStack, Flex, Select
} from "@chakra-ui/react";
import { Routes, Route } from "react-router-dom";
import { getKittens, getAccessToken, logout, refreshToken } from "./api"; 
import KittenCard from "./components/KittenCard";
import AddKittenForm from "./components/AddKittenForm";
import Login from "./components/Login"; 
import KittenPage from "./components/KittenPage";

export default function App() {
  const [kittens, setKittens] = useState([]);
  const [loading, setLoading] = useState(true);
  const [page, setPage] = useState(1);
  
  const [breed, setBreed] = useState("");
  const [sort, setSort] = useState("");

  const [isAuthenticated, setIsAuthenticated] = useState(!!getAccessToken());
  const [isCheckingAuth, setIsCheckingAuth] = useState(true);
  const [userRole, setUserRole] = useState(localStorage.getItem("role") || "guest");
  
  const [showLogin, setShowLogin] = useState(false);
  const toast = useToast();

  useEffect(() => {
    const checkAuth = async () => {
      const token = getAccessToken();
      if (!token) {
        try {
          const data = await refreshToken(); 
          if (data.access_token) setIsAuthenticated(true);
        } catch (err) {
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
      const res = await getKittens(page, 6, breed, sort);
      setKittens(res.data);
    } catch (err) {
      toast({ title: "Не удалось загрузить котят", status: "error" });
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    load();
  }, [page, breed, sort]);

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

  if (isCheckingAuth) return <Center minH="100vh"><Spinner size="xl" color="purple.500" /></Center>;

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
    <Routes>
      <Route path="/" element={
        <Box minH="100vh" bgGradient="linear(to-b, purple.50, white)">
          <Box bg="purple.600" color="white" py={4} mb={8} shadow="md">
            <Container maxW="container.xl">
              <Flex justify="space-between" align="center">
                <HStack spacing={3}>
                  <Text fontSize="3xl">🐱</Text>
                  <Box>
                    <Heading size="md" letterSpacing="wide">Kittens Shop · EU</Heading>
                  </Box>
                </HStack>
                {isAuthenticated ? (
                  <Button colorScheme="whiteAlpha" size="sm" onClick={handleLogout}>Выйти</Button>
                ) : (
                  <Button colorScheme="whiteAlpha" size="sm" onClick={() => setShowLogin(true)}>Войти</Button>
                )}
              </Flex>
            </Container>
          </Box>

          <Container maxW="container.xl" pb={16}>
            <Box 
              mb={8} p={8} borderRadius="2xl" 
              bgGradient="linear(to-r, purple.500, pink.400)" 
              color="white" shadow="xl"
            >
              <Heading size="xl" mb={2}>Найди своего идеального друга</Heading>
              <Text fontSize="lg" opacity={0.9}>Более 1000 счастливых котиков уже нашли свой дом</Text>
            </Box>

            <Flex mb={6} gap={4} wrap="wrap" bg="white" p={4} borderRadius="lg" shadow="sm">
              <Select placeholder="Все породы" value={breed} onChange={(e) => { setBreed(e.target.value); setPage(1)}} maxW="200px" bg="gray.50">
                <option value="Британская">Британская</option>
                <option value="Шотландская">Шотландская</option>
                <option value="Мейн-кун">Мейн-кун</option>
              </Select>
              <Select placeholder="Сортировка" value={sort} onChange={(e) => { setSort(e.target.value); setPage(1); }} maxW="200px" bg="gray.50">
                <option value="price_asc">Сначала дешевые</option>
                <option value="price_desc">Сначала дорогие</option>
              </Select>
            </Flex>

            <Tabs colorScheme="purple" variant="soft-rounded">
              <TabList mb={4}>
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
                              <KittenCard key={k.ID} kitten={k} />
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

          <Box bg="gray.800" color="gray.400" py={8} mt={16}>
            <Container maxW="container.xl">
              <Flex justify="space-between" align="center">
                <Text>© 2026 Kittens Shop. Все права защищены.</Text>
              </Flex>
            </Container>
          </Box>
        </Box>
      } />
      <Route path="/kitten/:id" element={<KittenPage />} />
    </Routes>
  );
}
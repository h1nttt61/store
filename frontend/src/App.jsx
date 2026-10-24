import { useEffect, useState } from "react";
import {
  Box, Container, Heading, SimpleGrid, Tabs, TabList, Tab, TabPanels, TabPanel,
  Spinner, Center, Text, useToast, Button, HStack, Flex
} from "@chakra-ui/react";
import { getKittens } from "./api";
import KittenCard from "./components/KittenCard";
import AddKittenForm from "./components/AddKittenForm";

export default function App() {
  const [kittens, setKittens] = useState([]);
  const [loading, setLoading] = useState(true);
  const [page, setPage] = useState(1);
  const toast = useToast();

  const load = async () => {
    setLoading(true);
    try {
      const res = await getKittens(page, 6);
      setKittens(res.data);
    } catch (err) {
      console.error(err);
      toast({
        title: "Не удалось загрузить котят",
        description: "Бэкенд запущен? Порт 3000?",
        status: "error",
        duration: 5000,
        isClosable: true,
      });
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    load();
  }, [page]);

  return (
      <Box minH="100vh" bg="gray.50">
        <Box bg="purple.600" color="white" py={6} mb={8}>
          <Container maxW="container.xl">
            <Heading size="lg">🐱 Kittens Shop · EU</Heading>
            <Text opacity={0.85} mt={1}>
              Твой проект по продаже котят
            </Text>
          </Container>
        </Box>

        <Container maxW="container.xl" pb={16}>
          <Tabs colorScheme="purple" variant="enclosed">
            <TabList>
              <Tab>Каталог</Tab>
              <Tab>Добавить</Tab>
            </TabList>

            <TabPanels>
              <TabPanel px={0}>
                {loading ? (
                    <Center py={20}>
                      <Spinner size="xl" color="purple.500" />
                    </Center>
                ) : kittens.length === 0 ? (
                    <Text textAlign="center" py={20} color="gray.500">
                      Пока нет доступных котят
                    </Text>
                ) : (
                    <>
                      <SimpleGrid columns={{ base: 1, md: 2, lg: 3 }} spacing={6}>
                        {kittens.map((k) => (
                            <KittenCard
                                key={k.ID}
                                kitten={k}
                                onDetails={(kitten) =>
                                    toast({
                                      title: kitten.name,
                                      description: kitten.description || "Без описания",
                                      status: "info",
                                      duration: 4000,
                                      isClosable: true,
                                    })
                                }
                            />
                        ))}
                      </SimpleGrid>

                      <Flex justify="center" mt={8}>
                        <HStack>
                          <Button
                              onClick={() => setPage(page - 1)}
                              isDisabled={page === 1}
                          >
                            Назад
                          </Button>
                          <Text fontWeight="bold" alignSelf="center" mx={4}>
                            Страница {page}
                          </Text>
                          <Button
                              onClick={() => setPage(page + 1)}
                              isDisabled={kittens.length < 6}
                          >
                            Вперед
                          </Button>
                        </HStack>
                       </Flex>
                    </>
                )}
              </TabPanel>

              <TabPanel px={0}>
                <AddKittenForm onCreated={() => { setPage(1); load(); }} />
              </TabPanel>
            </TabPanels>
          </Tabs>
        </Container>
      </Box>
  );
}
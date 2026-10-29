import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { Box, Button, Heading, Text, Image, Container, Spinner, Center, Badge } from "@chakra-ui/react";
import { getKitten } from "../api";

export default function KittenPage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [kitten, setKitten] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    getKitten(id)
      .then((data) => setKitten(data))
      .catch((err) => console.error(err))
      .finally(() => setLoading(false));
  }, [id]);

  if (loading) return <Center minH="100vh"><Spinner size="xl" color="purple.500" /></Center>;
  if (!kitten) return <Center minH="100vh"><Text>Котик не найден 😿</Text></Center>;

  return (
    <Container maxW="container.md" py={10}>
      <Button onClick={() => navigate(-1)} mb={6}>← Назад</Button>
      
      <Box bg="white" borderRadius="lg" overflow="hidden" shadow="xl">
        <Image src={kitten.image_url || "https://placekitten.com/800/400"} alt={kitten.name} w="100%" h="400px" objectFit="cover" />
        <Box p={8}>
          <Heading mb={2}>{kitten.name}</Heading>
          <Badge colorScheme={kitten.is_available ? "green" : "red"} mb={4}>
            {kitten.is_available ? "Доступен" : "Продан"}
          </Badge>
          <Text fontSize="xl" fontWeight="bold" color="purple.600" mb={6}>€{kitten.price}</Text>
          
          <Text color="gray.600" mb={4}>Порода: {kitten.breed}</Text>
          <Text color="gray.600" mb={6}>Возраст: {kitten.age_months} мес.</Text>
          
          <Text fontSize="lg" mb={8}>{kitten.description || "Ласковый, игривый котик ищет новый дом."}</Text>
          
          <Button colorScheme="purple" size="lg" width="100%" isDisabled={!kitten.is_available}>
            {kitten.is_available ? "Купить" : "Продан"}
          </Button>
        </Box>
      </Box>
    </Container>
  );
}
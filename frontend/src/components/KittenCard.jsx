import { Box, Image, Badge, Heading, Text, HStack, Button } from "@chakra-ui/react";
import { Link } from "react-router-dom";

export default function KittenCard({ kitten }) {
  return (
    <Box
      borderWidth="1px"
      borderRadius="lg"
      overflow="hidden"
      bg="white"
      shadow="md"
      _hover={{ shadow: "xl", transform: "translateY(-2px)" }}
      transition="all 0.2s"
    >
      <Image
        src={kitten.image_url || "https://placekitten.com/400/300"}
        alt={kitten.name}
        h="200px"
        w="100%"
        objectFit="cover"
      />

      <Box p={4}>
        <HStack justify="space-between" mb={2}>
          <Heading size="md">{kitten.name}</Heading>
          <Badge colorScheme={kitten.is_available ? "green" : "red"}>
            {kitten.is_available ? "Доступен" : "Продан"}
          </Badge>
        </HStack>

        <Text fontSize="sm" color="gray.600" mb={1}>
          Порода: {kitten.breed}
        </Text>
        <Text fontSize="sm" color="gray.600" mb={1}>
          Возраст: {kitten.age_months} мес.
        </Text>
        <Text fontSize="lg" fontWeight="bold" color="purple.600" mb={3}>
          €{kitten.price}
        </Text>

        <Button as={Link} to={`/kitten/${kitten.ID}`} colorScheme="purple" size="sm" width="100%">
          Подробнее
        </Button>
      </Box>
    </Box>
  );
}
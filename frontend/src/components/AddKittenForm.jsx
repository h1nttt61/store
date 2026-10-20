import { useState } from "react";
import { Box, Button, FormControl, FormLabel, Input, NumberInput, NumberInputField, Textarea, VStack, useToast } from "@chakra-ui/react";
import { createKitten } from "../api";

const empty = {
    name: "",
    breed: "",
    age_months: 0,
    price: 0,
    description: "",
    image_url: "",
}

export default function AddKittenForm({onCreated}){
    const [form, setForm] = useState(empty);
    const [loading, setLoading] = useState(false);
    const toast = useToast();

    const set = (field) => (e) => setForm({...form, [field]: e.target.value});

    const handleSubmit = async (e) =>{
        e.preventDefault();
        setLoading(true);
        try{
            const payload ={
                ...form,
                age_months: Number(form.age_months),
                price: Number(form.price)
            };
            const created = await createKitten(payload);
            toast({
                title: "Котенок добавлен",
                status: "success",
                duration: 3000,
                isClosable: true,
            });
            setForm(empty);
            onCreated(created);
        }
        catch(err){
            console.error(err);
             toast({
                title: "Ошибка при создании",
                description: "Ошибка",
                status: "error",
                duration: 4000,
                isClosable: true,
            });
        }finally{
            setLoading(false);
        }
    };
    return (
        <Box as="form" onSubmit={handleSubmit} bg="white" p={6} borderRadius="lg" shadow="md">
            <VStack spacing={4} align="stretch">
                <FormControl isRequired>
                    <FormLabel>Имя</FormLabel>
                    <Input value={form.name} onChange={set("name")} placeholder="Мурзик"></Input>
                </FormControl>
                <FormControl isRequired>
                    <FormLabel>Порода</FormLabel>
                    <Input value={form.breed} onChange={set("breed")} placeholder="Британская"></Input>
                </FormControl>
                <FormControl isRequired>
                    <FormLabel>Возраст (мес.)</FormLabel>
                    <NumberInput min={0}>
                        <NumberInputField value={form.age_months} onChange={set("age_months")}></NumberInputField>
                    </NumberInput>
                </FormControl>
                <FormControl isRequired>
                    <FormLabel>Цена ($)</FormLabel>
                    <NumberInput min={0}>
                        <NumberInputField value={form.price} onChange={set("price")}></NumberInputField>
                    </NumberInput>
                </FormControl>
                <FormControl>
                    <FormLabel>Описание</FormLabel>
                    <Textarea value={form.description} onChange={set("description")} placeholder="Ласковый, игривый"></Textarea>
                </FormControl>
                <FormControl>
                    <FormLabel>URL фотo</FormLabel>
                    <Input value={form.image_url} onChange={set("image_url")} placeholder="https://..."></Input>
                </FormControl>
                <Button type="submit" colorScheme="purple" isLoading={loading}>
                    Добавить
                </Button>
            </VStack>
        </Box>
    )
}
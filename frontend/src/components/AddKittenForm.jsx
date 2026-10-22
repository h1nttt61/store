import { useState } from "react";
import { Box, Button, FormControl, FormLabel, Input, NumberInput, NumberInputField, Textarea, VStack, useToast, FormErrorMessage  } from "@chakra-ui/react";
import { createKitten, uploadImage } from "../api";

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
    const [errors, setErrors] = useState({});
    const [file, setFile] = useState(null);
    const toast = useToast();

    const set = (field) => (e) => setForm({...form, [field]: e.target.value});

    const handleSubmit = async (e) =>{
        e.preventDefault();
        setLoading(true);
        setErrors({});

      
        try{


            let finalImageUrl = form.image_url

            if (file){
                const uploadRes = await uploadImage(file);
                finalImageUrl = `http://localhost:3000${uploadRes.url}`; 
            }

            const payload ={
                ...form,
                image_url: finalImageUrl,
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

            if (err.response && err.response.status === 400 && err.response.data.details){
             setErrors(err.response.data.details); 
             toast({
                title: "Ошибка валидации",
                description: "Проверьте правильность заполнения полей",
                status: "error",
                duration: 4000,
                isClosable: true,
            });
        }else{
            toast({
                title: "Ошибка при создании",
                description: "Ошибка",
                status: "error",
                duration: 4000,
                isClosable: true,
            });
        }
        }finally{
            setLoading(false);
        }
    };
    return (
        <Box as="form" onSubmit={handleSubmit} bg="white" p={6} borderRadius="lg" shadow="md">
            <VStack spacing={4} align="stretch">
                <FormControl isInvalid={!!errors.name} isRequired>
                    <FormLabel>Имя</FormLabel>
                    <Input value={form.name} onChange={(e) => {set("name")(e); if (errors.name) setErrors({...errors, name: null})}} placeholder="Мурзик"></Input>
                    {errors.name && <FormErrorMessage>{errors.name}</FormErrorMessage>}
                </FormControl>
                <FormControl isInvalid={!!errors.breed} isRequired>
                    <FormLabel>Порода</FormLabel>
                    <Input value={form.breed} onChange={(e) => { set("breed")(e); if (errors.breed) setErrors({...errors, breed: null}); }}  placeholder="Британская"></Input>
                    {errors.breed && <FormErrorMessage>{errors.breed}</FormErrorMessage>}
                </FormControl>
                <FormControl isRequired>
                    <FormLabel>Возраст (мес.)</FormLabel>
                    <NumberInput min={0}>
                        <NumberInputField value={form.age_months} onChange={set("age_months")}></NumberInputField>
                    </NumberInput>
                </FormControl>
                <FormControl  isInvalid={!!errors.price} isRequired>
                    <FormLabel>Цена ($)</FormLabel>
                    <NumberInput min={0}>
                        <NumberInputField value={form.price} onChange={(e) => { set("price")(e); if (errors.price) setErrors({...errors, price: null});}}></NumberInputField>
                    </NumberInput>
                    {errors.price && <FormErrorMessage>{errors.price}</FormErrorMessage>}
                </FormControl>
                <FormControl>
                    <FormLabel>Описание</FormLabel>
                    <Textarea value={form.description} onChange={set("description")} placeholder="Ласковый, игривый"></Textarea>
                </FormControl>
                 <FormControl>
                    <FormLabel>Фото котенка</FormLabel>
                    <Input type="file" accept="image/*" onChange={(e) => setFile(e.target.files[0])} p={1}></Input>
                </FormControl>
                <Button type="submit" colorScheme="purple" isLoading={loading}>
                    Добавить
                </Button>
            </VStack>
        </Box>
    )
}
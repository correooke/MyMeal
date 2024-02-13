"use client";

import Box from "@mui/material/Box";
import ProductForm from "../ProductForm/ProductForm";
import { Product } from "@/models/product";
import { useTheme } from "@mui/material/styles";
import CategoryInput from "./components/CategoryInput";
import AddNewProductButton from "./components/AddNewProductButton";
import AddNewCategoryButton from "./components/AddNewCategoryButton";
import { Typography } from "@mui/material";
import Actions from "../Shared/Actions";
import { use, useState } from "react";

const mockProduct: Product = {
  name: "",
  price: 10.99,
  description: "Esta es una descripción de ejemplo para el producto.",
  id: 0,
  imageUrl: "",
  isActive: false,
};

const CategoryForm = () => {
  const theme = useTheme();
  const [products, setProducts] = useState([mockProduct]);
  return (
    <Box
      display={"flex"}
      sx={{
        height: "100%",
        paddingTop: "10px",
        width: theme.mainWidth,
        flexDirection: "column",
        gap: "10px",
        boxSizing: "border-box",
      }}
    >
      <Box
        sx={{
          width: "100%",
          padding: 0,
          margin: 0,
          justifyContent: "center",
          display: "flex",
          gap: "5px",
          boxSizing: "border-box",
          paddingRight: "25px",
        }}
      >
        <Box
          sx={{
            display: "flex",
            boxSizing: "border-box",
            paddingLeft: "25px",
            paddingRight: "25px",
            justifyContent: "center",
            flexGrow: 1,
            backgroundColor: theme.palette.primary.main,
            borderRadius: "4px",
          }}
        >
          <Typography
            variant="h5"
            color="white"
            sx={{ lineHeight: "56px", height: "56px", textAlign: "center" }}
          >
            01
          </Typography>
        </Box>
        <CategoryInput
          value={""}
          onChange={function (newValue: string, name: string): void {}}
        ></CategoryInput>
        <Actions />
      </Box>
      <Box
        sx={{
          display: "flex",
          gap: "10px",
          flexDirection: "column",
          alignItems: "flex-end",
        }}
      >
        {products.map((product) => (
          <ProductForm key={product.id} product={mockProduct} />
        ))}

        <Box sx={{ flexGrow: 1, justifyContent: "flex-end" }}>
          <AddNewProductButton
            onClick={() => {
              setProducts([mockProduct, ...products]);
            }}
          />
        </Box>
      </Box>
      <Box
        sx={{
          width: "100%",
          display: "flex",
          gap: "20px",
          justifyContent: "flex-end",
          alignItems: "flex-end",
          flexDirection: "column",
          borderTop: `5px solid ${theme.palette.primary.main}`,
        }}
      >
        <Box sx={{ flexGrow: 1, marginTop: "10px" }}>
          <AddNewCategoryButton />
        </Box>
      </Box>
    </Box>
  );
};

export default CategoryForm;

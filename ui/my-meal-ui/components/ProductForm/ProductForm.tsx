"use client";

import { Product } from "@/models/product";
import { Box, Grid } from "@mui/material";
import { FC, useState } from "react";
import ProductInput from "./components/ProductInput";
import ProductDescriptionInput from "./components/ProductDescriptionInput";
import PriceInput from "./components/PriceInput";
import DeleteButton from "../Shared/DeleteButton";
import IsActiveOption from "../Shared/IsActiveOption";
import ProductDetails from "./components/ProductDetails";
import ProductImageUploader from "./components/ProductImageUploader";
import { useTheme } from "@mui/material/styles";
import Actions from "../Shared/Actions";

interface ProductFormProps {
  product: Product;
}

const ProductForm: FC<ProductFormProps> = ({ product }) => {
  const [updatedProduct, setProduct] = useState({
    ...product,
  });

  const handleChange = (value: string | number, name: string) => {
    setProduct({
      ...updatedProduct,
      [name]: value,
    });
  };

  const theme = useTheme();

  return (
    <Box
      display={"flex"}
      sx={{
        width: theme.mainWidth,
        padding: "25px",
        boxSizing: "border-box",
        borderRadius: theme.border.radius,
        border: `1px solid ${theme.palette.primary.main}`,
        gap: "10px",
      }}
    >
      <Box>
        <ProductImageUploader />
      </Box>
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          width: "100%",
          gap: "10px",
        }}
      >
        <Box
          sx={{
            width: "100%",
            display: "flex",
            flexDirection: "row",
          }}
        >
          <ProductInput value={updatedProduct.name} onChange={handleChange} />
        </Box>

        <ProductDescriptionInput
          value={updatedProduct.description}
          onChange={handleChange}
        />
        <Box sx={{ display: "flex", gap: "20px" }}>
          <Grid item xs={6}>
            <ProductDetails product={updatedProduct} />{" "}
          </Grid>
          <Grid item xs={6}>
            <PriceInput value={updatedProduct.price} onChange={handleChange} />
          </Grid>
        </Box>
        <Box sx={{ display: "flex", justifyContent: "flex-end" }}>
          <Actions></Actions>
        </Box>
      </Box>
    </Box>
  );
};

export default ProductForm;

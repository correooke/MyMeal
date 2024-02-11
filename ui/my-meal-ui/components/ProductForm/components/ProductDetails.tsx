import { Box, Typography } from "@mui/material";
import { Product } from "@/models/product";

type ProductDetailsProps = {
  product: Product;
};

const ProductDetails = ({ product }: ProductDetailsProps) => {
  return (
    <Box border={1} padding={2}>
      <Typography variant="body1">Detalles</Typography>
    </Box>
  );
};

export default ProductDetails;

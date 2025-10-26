import { buildClientSchema, printSchema } from "graphql";
import fs from "fs";

const [,, inputFile, outputFile] = process.argv;
const introspection = JSON.parse(fs.readFileSync(inputFile, "utf-8"));

const schema = buildClientSchema(introspection);
fs.writeFileSync(outputFile, printSchema(schema));
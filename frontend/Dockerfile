# frontend/Dockerfile
FROM node:16-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем package.json и устанавливаем зависимости
COPY package*.json ./
RUN npm install

# Копируем все остальные файлы
COPY . .

# Стартуем приложение
CMD ["npm", "start"]

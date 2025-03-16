FROM nginx:alpine

# Copy static files
COPY index.html /usr/share/nginx/html/
COPY css/styles.css /usr/share/nginx/html/
COPY js/app.js /usr/share/nginx/html/

EXPOSE 80

# Use default nginx configuration
CMD ["nginx", "-g", "daemon off;"]
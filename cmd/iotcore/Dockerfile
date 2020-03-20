FROM arm32v7/node

ENV IOT_KEY_PATH=test_device.private.key
ENV IOT_CERT_PATH=test_device.cert.pem
ENV IOT_CLIENT_ID=a2vd0mrnqc6v51-ats.iot.eu-west-1.amazonaws.com
ENV IOT_HOST=sdk-nodejs-349708a4-3783-4bbd-bf68-42dc098df24c

WORKDIR /usr/src/app

COPY package.json .
COPY package-lock.json .

RUN npm install

RUN curl https://www.amazontrust.com/repository/AmazonRootCA1.pem > root-CA.crt

# Bundle app source
COPY . .

CMD [ "node", "index.js" ]
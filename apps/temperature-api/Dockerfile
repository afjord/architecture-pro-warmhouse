FROM gradle:9.2.1-jdk25 as builder
COPY /app /app
WORKDIR /app
RUN gradle clean build

FROM amazoncorretto:25-jdk
COPY --from=builder app/build/libs/app.jar app.jar
EXPOSE 8081
CMD ["java", "-jar", "app.jar"]
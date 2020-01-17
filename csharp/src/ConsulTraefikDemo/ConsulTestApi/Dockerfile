FROM microsoft/dotnet:2.2-aspnetcore-runtime AS base
WORKDIR /app
EXPOSE 80
RUN mkdir /logs

FROM microsoft/dotnet:2.2-sdk AS restore
WORKDIR /build

COPY *.csproj ./src/
RUN dotnet restore /build/src

FROM restore AS source
WORKDIR /build
ADD . ./src/

FROM source AS publish
WORKDIR /build/src
RUN dotnet publish -c Release -o /app

FROM base AS final
WORKDIR /app
COPY --from=publish /app .
ENTRYPOINT ["dotnet", "ConsulTestApi.dll"]
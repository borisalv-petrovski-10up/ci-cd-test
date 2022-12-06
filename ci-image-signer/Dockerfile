FROM gcr.io/sk_borislav/kritis-signer as kritis

FROM gcr.io/cloud-builders/gcloud:latest as runtime

COPY --from=kritis /kritis/signer /kritis/signer

COPY sign-image.sh .
RUN chmod +x sign-image.sh 

ENTRYPOINT ["bash"]
CMD ["/sign-image.sh"]

FROM rabbitmq:3-management

RUN rabbitmq-plugins enable --offline rabbitmq_mqtt

RUN rabbitmq-plugins enable --offline rabbitmq_federation_management 

RUN rabbitmq-plugins enable --offline rabbitmq_stomp

RUN rabbitmq-plugins enable --offline rabbitmq_consistent_hash_exchange
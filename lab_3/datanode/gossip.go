package main

import (
	"context"
	"log"
	"time"

	cmpb "lab_3/proto/common/cmpb"
	dpb "lab_3/proto/datanode/dpb"
)

func (s *DatanodeServer) sendGossip(update *cmpb.FlightUpdate) {
	for _, peer := range s.gossipAddrs {
		s.mu.Lock()
		cli := s.gossipClients[peer]
		s.mu.Unlock()

		if cli == nil {
			cli = s.dialPeer(peer)
			if cli == nil {
				continue
			}

			s.mu.Lock()
			s.gossipClients[peer] = cli
			s.mu.Unlock()
		}

		go func(cli dpb.DatanodeServiceClient, peerAddr string) {
			_, err := cli.Gossip(context.Background(), update)
			if err != nil {
				log.Printf("[Datanode %s] Gossip error hacia %s: %v", s.id, peerAddr, err)
				s.mu.Lock()
				s.gossipClients[peerAddr] = nil
				s.mu.Unlock()
			}
		}(cli, peer)
	}
}

func (s *DatanodeServer) StartPeriodicGossip() {
	ticker := time.NewTicker(3 * time.Second)

	log.Printf("[Datanode %s] Gossip periódico activado.", s.id)

	for {
		<-ticker.C

		s.mu.Lock()
		for flightID, st := range s.flights {
			update := &cmpb.FlightUpdate{
				FlightId:    flightID,
				Status:      st.Status,
				Gate:        st.Gate,
				VectorClock: copyVectorClock(st.VectorClock),
			}
			go s.sendGossip(update)
		}
		s.mu.Unlock()
	}
}

func (s *DatanodeServer) Gossip(ctx context.Context, update *cmpb.FlightUpdate) (*cmpb.Empty, error) {

	s.applyUpdateLocal(update)
	return &cmpb.Empty{}, nil
}

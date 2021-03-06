package metainfo

import "github.com/anacrolix/missinggo"

type Piece struct {
	Info *InfoEx
	i    int
}

func (p Piece) Length() int64 {
	if p.i == p.Info.NumPieces()-1 {
		return p.Info.TotalLength() - int64(p.i)*p.Info.PieceLength
	}
	return p.Info.PieceLength
}

func (p Piece) Offset() int64 {
	return int64(p.i) * p.Info.PieceLength
}

func (p Piece) Hash() (ret Hash) {
	missinggo.CopyExact(&ret, p.Info.Pieces[p.i*20:(p.i+1)*20])
	return
}

func (p Piece) Index() int {
	return p.i
}

func (p Piece) Key() PieceKey {
	return PieceKey{p.Info.Hash(), p.i}
}

type PieceKey struct {
	Hash  Hash
	Index int
}

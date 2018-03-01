package goseriald

import (
	"github.com/albenik/iolog"
	"go.bug.st/serial.v1"
)

type PortWrapper struct {
	port  serial.Port
	iolog *iolog.IOLog
}

func Open(name string, mode *serial.Mode, log *iolog.IOLog) (wr *PortWrapper, outerr error) {
	log.LogAny("open", func() (interface{}, error) {
		if port, err := serial.Open(name, mode); err == nil {
			wr = &PortWrapper{port: port, iolog: log}
			return mode, nil
		} else {
			outerr = err
			return mode, err
		}
	})
	return
}

func (pw *PortWrapper) GetName() string {
	return pw.port.GetName()
}

func (pw *PortWrapper) SetMode(mode *serial.Mode) error {
	return pw.iolog.LogAny("set_mode", func() (interface{}, error) {
		err := pw.port.SetMode(mode)
		return mode, err
	})
}

func (pw *PortWrapper) SetReadTimeout(t int) error {
	return pw.iolog.LogAny("set_read_timeout", func() (interface{}, error) {
		err := pw.port.SetReadTimeout(t)
		return t, err
	})
}

func (pw *PortWrapper) SetReadTimeoutEx(t, i uint32) error {
	return pw.iolog.LogAny("set_read_timeout_ex", func() (interface{}, error) {
		err := pw.port.SetReadTimeoutEx(t, i)
		return &struct{ T, I uint32 }{t, i}, err
	})
}

func (pw *PortWrapper) SetFirstByteReadTimeout(t uint32) error {
	return pw.iolog.LogAny("set_first_byte_read_timeout", func() (interface{}, error) {
		err := pw.port.SetFirstByteReadTimeout(t)
		return t, err
	})
}

func (pw *PortWrapper) SetWriteTimeout(t int) error {
	return pw.iolog.LogAny("set_write_timeout", func() (interface{}, error) {
		err := pw.port.SetWriteTimeout(t)
		return t, err
	})
}

func (pw *PortWrapper) ReadyToRead() (r uint32, err error) {
	pw.iolog.LogAny("ready_to_read", func() (interface{}, error) {
		r, err = pw.port.ReadyToRead()
		return r, err
	})
	return
}

func (pw *PortWrapper) Read(p []byte) (n int, err error) {
	return pw.iolog.LogIO("read", pw.port.Read, p)
}

func (pw *PortWrapper) Write(p []byte) (n int, err error) {
	return pw.iolog.LogIO("write", pw.port.Write, p)
}

func (pw *PortWrapper) ResetInputBuffer() error {
	return pw.iolog.LogAny("reset_input_buffer", func() (interface{}, error) {
		err := pw.port.ResetInputBuffer()
		return nil, err
	})
}

func (pw *PortWrapper) ResetOutputBuffer() error {
	return pw.iolog.LogAny("reset_output_buffer", func() (interface{}, error) {
		err := pw.port.ResetOutputBuffer()
		return nil, err
	})
}

func (pw *PortWrapper) SetDTR(dtr bool) error {
	return pw.iolog.LogAny("set_dtr", func() (interface{}, error) {
		err := pw.port.SetDTR(dtr)
		return dtr, err
	})
}

func (pw *PortWrapper) SetRTS(rts bool) error {
	return pw.iolog.LogAny("set_rts", func() (interface{}, error) {
		err := pw.port.SetRTS(rts)
		return rts, err
	})
}

func (pw *PortWrapper) GetModemStatusBits() (bits *serial.ModemStatusBits, err error) {
	pw.iolog.LogAny("get_modem_status_bits", func() (interface{}, error) {
		bits, err = pw.port.GetModemStatusBits()
		return bits, err
	})
	return
}

func (pw *PortWrapper) Close() error {
	return pw.iolog.LogAny("close", func() (interface{}, error) {
		err := pw.port.Close()
		return nil, err
	})
}

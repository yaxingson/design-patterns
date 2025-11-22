package main

type Storage interface {
	set()
	get()
	remove()
}

type MemoryStorage struct{}

func (m *MemoryStorage) set()    {}
func (m *MemoryStorage) get()    {}
func (m *MemoryStorage) remove() {}

type FileStorage struct{}

func (f *FileStorage) set()    {}
func (f *FileStorage) get()    {}
func (f *FileStorage) remove() {}

type DbStorage struct{}

func (d *DbStorage) set()    {}
func (d *DbStorage) get()    {}
func (d *DbStorage) remove() {}

type CryptoStorage struct {
	storage Storage
}

func (c *CryptoStorage) set() {
	// extra operate
	c.storage.set()
}
func (c *CryptoStorage) get()    {}
func (c *CryptoStorage) remove() {}

func init() {

}

/* -----------------------------------------------------------------------------
 * std_map.i
 *
 * SWIG typemaps for std::map
 * ----------------------------------------------------------------------------- */

%include <std_common.i>

// ------------------------------------------------------------------------
// std::map
// ------------------------------------------------------------------------

%{
#include <map>
#include <algorithm>
#include <stdexcept>
#include <vector>
%}

// exported class

namespace std {

    template<class K, class T, class C = std::less<K> > class map {
        // add typemaps here
      public:
        typedef size_t size_type;
        typedef ptrdiff_t difference_type;
        typedef K key_type;
        typedef T mapped_type;
        typedef std::pair< const K, T > value_type;
        typedef value_type* pointer;
        typedef const value_type* const_pointer;
        typedef value_type& reference;
        typedef const value_type& const_reference;

        map();
        map(const map& other);
        
        unsigned int size() const;
        bool empty() const;
        void clear();
        %extend {
            const T& get(const K& key) throw (std::out_of_range) {
                std::map< K, T, C >::iterator i = self->find(key);
                if (i != self->end())
                    return i->second;
                else
                    throw std::out_of_range("key not found");
            }
            void set(const K& key, const T& x) {
                (*self)[key] = x;
            }
            void del(const K& key) throw (std::out_of_range) {
                std::map< K, T, C >::iterator i = self->find(key);
                if (i != self->end())
                    self->erase(i);
                else
                    throw std::out_of_range("key not found");
            }
            bool has_key(const K& key) {
                std::map< K, T, C >::iterator i = self->find(key);
                return i != self->end();
            }
            std::vector<K> keys() {
                std::vector<K> keys;
                keys.reserve(self->size());
                for (std::map<  K, T, C >::iterator it = self->begin(); it != self->end(); ++it) {
                    keys.push_back(it->first);
                }
                return keys;
            }
            std::vector<T> values() {
                std::vector<T> values;
                values.reserve(self->size());
                for (std::map<  K, T, C >::iterator it = self->begin(); it != self->end(); ++it) {
                    values.push_back(it->second);
                }
                return values;
            }
        }
    };
}